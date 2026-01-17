package notify

type Sender interface {
	Send(body string) error
}

func SendMessage(S Sender, body string) error {
	return S.Send(body)

}
