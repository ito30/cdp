package notifier

import "fmt"

type whatsAppNotifier struct {
	notifier INotifier
}

func NewWhatsAppNotifier(notifier INotifier) *whatsAppNotifier {
	return &whatsAppNotifier{
		notifier: notifier,
	}
}

func (n *whatsAppNotifier) Send(message string) error {
	fmt.Println("--sending notification to whatsApp--")

	n.notifier.Send(message)
	return nil
}
