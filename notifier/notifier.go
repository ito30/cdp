package notifier

import "fmt"

type INotifier interface {
	Send(message string) error
}

type notifier struct{}

func (n *notifier) Send(message string) error {
	fmt.Println(message)
	return nil
}
