// Package rsa provides RSA encryption, decryption, signing and verification with hashing
package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// PrivateKey ...
var PrivateKey *rsa.PrivateKey

// HashSum ...
var HashSum []byte

// GeneratePrivateKey ...
func GeneratePrivateKey(privateKey *rsa.PrivateKey) {
	if privateKey != nil {
		PrivateKey = privateKey
		return
	}

	var err error
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic(err)
	}
}

// EncryptSHA256 ...
func EncryptSHA256(content []byte) ([]byte, error) {
	publicKey := PrivateKey.PublicKey
	return rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		content,
		nil)
}

// DecryptSHA256 ...
func DecryptSHA256(encryptedBytes []byte) ([]byte, error) {
	return PrivateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

// SignSHA256 ...
func SignSHA256() ([]byte, error) {
	return rsa.SignPSS(rand.Reader, PrivateKey, crypto.SHA256, HashSum, nil)
}

// VerifySHA256 ...
func VerifySHA256(signature []byte) error {
	publicKey := PrivateKey.PublicKey
	return rsa.VerifyPSS(&publicKey, crypto.SHA256, HashSum, signature, nil)
}

// GenerateHash256Sum ...
func GenerateHash256Sum(message []byte) {
	var err error
	msgHash := sha256.New()
	_, err = msgHash.Write(message)

	if err != nil {
		panic(err)
	}

	HashSum = msgHash.Sum(nil)
}
