package notify

import "fmt"

type FakeSender struct {
}

func (f *FakeSender) Send(body string, errchan chan error) {
	fmt.Println("fake sending", body)

}
