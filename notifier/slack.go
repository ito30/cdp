package notifier

import "fmt"

type slackNotifier struct {
	notifier Notifier
}

func NewSlackNotifier(notifier Notifier) *slackNotifier {
	return &slackNotifier{
		notifier: notifier,
	}
}

func (n *slackNotifier) Send(message string) error {
	fmt.Println("--sending notification to slack--")

	n.notifier.Send(message)
	return nil
}
