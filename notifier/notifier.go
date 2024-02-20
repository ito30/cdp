package notifier

import "fmt"

type INotifier interface {
	Send(message string) error
}

type (
	NotifierType  string
	NotifierOpt   map[NotifierType]bool
	NotifierTypes map[NotifierType]func(INotifier) INotifier
)

var (
	// NotifierType
	NotifierTypeSlack NotifierType = "slack"
	NotifierTypeSms   NotifierType = "sms"
	NotifierTypeWa    NotifierType = "wa"

	// Notifier instantiation
	notifiers NotifierTypes = NotifierTypes{
		NotifierTypeSlack: func(n INotifier) INotifier {
			return NewSlackNotifier(n)
		},
		NotifierTypeSms: func(n INotifier) INotifier {
			return NewSmsNotifier(n)
		},
		NotifierTypeWa: func(n INotifier) INotifier {
			return NewWhatsAppNotifier(n)
		},
	}
)

func NewN(opt NotifierOpt) INotifier {
	var n INotifier = &notifier{}

	for nType, active := range opt {
		if newN, ok := notifiers[nType]; ok && active {
			n = newN(n)
		}
	}

	return n
}

type notifier struct{}

func (n *notifier) Send(message string) error {
	fmt.Println(message)
	return nil
}
