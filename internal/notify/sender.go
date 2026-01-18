package notify

type Sender interface {
	Send(body string, errchan chan error)
}

func SendMessage(S Sender, body string, errchan chan error) {
	S.Send(body, errchan)

}
