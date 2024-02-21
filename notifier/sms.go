package notifier

import "fmt"

type smsNotifier struct {
	notifier Notifier
}

func NewSmsNotifier(notifier Notifier) *smsNotifier {
	return &smsNotifier{
		notifier: notifier,
	}
}

func (n *smsNotifier) Send(message string) error {
	fmt.Println("--sending notification to sms--")

	n.notifier.Send(message)
	return nil
}
