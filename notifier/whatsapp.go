package notifier

import "fmt"

type whatsAppNotifier struct {
	notifier Notifier
}

func NewWhatsAppNotifier(notifier Notifier) *whatsAppNotifier {
	return &whatsAppNotifier{
		notifier: notifier,
	}
}

func (n *whatsAppNotifier) Send(message string) error {
	fmt.Println("--sending notification to whatsApp--")

	n.notifier.Send(message)
	return nil
}
