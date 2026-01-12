package sms

type Sender interface {
	Send(body string) error
}

func SendMessage(S Sender, body string) error {
	return S.Send(body)

}
