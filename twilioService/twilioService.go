package twilioService

import (
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Client ...
var Client *twilio.RestClient
var twilioPhoneNumber string

// NewClient ...
func NewClient(accountSID, authToken, phoneNumber string) {
	Client = twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSID,
		Password: authToken,
	})
	twilioPhoneNumber = phoneNumber
}

// SendSMS ...
func SendSMS(messageBody, phoneNumber string) (*openapi.ApiV2010Message, error) {
	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(twilioPhoneNumber)
	params.SetBody(messageBody)
	return Client.ApiV2010.CreateMessage(params)
}
