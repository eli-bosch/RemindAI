package api

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	from   = "+1"
	client = twilio.NewRestClient()
)

func sendText(to string, body string) {
	params := &api.CreateMessageParams{}
	params.SetBody(body)
	params.SetFrom(from)
	params.SetTo(to)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Printf("Error sending text: %v", err)
	}
}

func listenText() {

}
