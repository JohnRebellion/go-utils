package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Client ...
var Client C

// C ...
type C struct {
	Client *http.Client
}

// New ...
func (c *C) New(client http.Client) {
	c.Client = &client
}

// RequestJSON ...
func RequestJSON(method, url string, input, output interface{}, headers http.Header) (*http.Response, error) {
	message, err := json.Marshal(input)

	if err == nil {
		headers.Set("Content-Type", "application/json")
		resp, err := Request(method, url, bytes.NewBuffer(message), headers)

		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				return resp, json.Unmarshal(body, &output)
			}
		}
	}

	return nil, err
}

// Request ...
func Request(method, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header = headers

	if err == nil {
		return Client.Client.Do(req)
	}

	return nil, err
}

// ReadBodyRequest ...
func ReadBodyRequest(method, url string, body io.Reader, headers http.Header) ([]byte, error) {
	resp, err := Request(method, url, body, headers)

	if err == nil {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}

	return nil, err
}
