package notify

import "fmt"

type FakeSender struct {
}

func (f *FakeSender) Send(body string) error {
	fmt.Println("fake sending", body)
	return nil

}
