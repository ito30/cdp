package notifier

import "fmt"

type smsNotifier struct {
	notifier INotifier
}

func NewSmsNotifier(notifier INotifier) *smsNotifier {
	return &smsNotifier{
		notifier: notifier,
	}
}

func (n *smsNotifier) Send(message string) error {
	fmt.Println("--sending notification to sms--")

	n.notifier.Send(message)
	return nil
}
