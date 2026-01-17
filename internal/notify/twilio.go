package notify

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioSender struct {
}

func (t *TwilioSender) Send(body string) error {
	from := os.Getenv("FromTwilioNumber")
	to := os.Getenv("ToTwilioNumber")
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

	_, err := client.Api.CreateMessage(params)
	// fmt.Println("Twilio SID:", *resp.Sid)
	// fmt.Println("Twilio Status:", *resp.Status)
	// fmt.Println("Twilio Error Code:", resp.ErrorCode)
	// fmt.Println("Twilio Error Message:", resp.ErrorMessage)

	return err

}
