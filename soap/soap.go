// Package soap provides SOAP XML Encoding and Decoding
package soap

import "encoding/xml"

// Envelope SOAP Envelope
type Envelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
	Header  Header
}

// Body SOAP Body
type Body struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

// Header SOAP Header
type Header struct {
	XMLName  struct{} `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

// Encode Encodes struct to XML SOAP Envelop
func Encode(contents, header interface{}) ([]byte, error) {
	data, err := xml.MarshalIndent(contents, "    ", "  ")

	if err != nil {
		return nil, err
	}

	headerData, err := xml.MarshalIndent(header, "    ", "  ")

	if err != nil {
		return nil, err
	}

	data = append([]byte("\n"), data...)
	envelope := Envelope{Body: Body{Contents: data}, Header: Header{Contents: headerData}}
	return xml.MarshalIndent(&envelope, "", "  ")
}

// Decode Decodes XML SOAP Envelop to struct
func Decode(data []byte, contents interface{}) error {
	envelope := Envelope{Body: Body{}}
	err := xml.Unmarshal(data, &envelope)

	if err != nil {
		return err
	}

	return xml.Unmarshal(envelope.Body.Contents, contents)
}
