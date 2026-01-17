package notify

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type WhatsAppSender struct {
}

func (w *WhatsAppSender) Send(body string) error {
	from := os.Getenv("From_Wapp_TwilioNumber")
	to := os.Getenv("To_Wapp_TwilioNumber")
	if from == "" || to == "" {
		return fmt.Errorf("Twilio numbers not set in env")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &twilioApi.CreateMessageParams{}

	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)
	fmt.Println(from, to, body)

	_, err := client.Api.CreateMessage(params)

	return err

}
