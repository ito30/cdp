package notifier

import "fmt"

type Notifier interface {
	Send(message string) error
}

type (
	NotifierType  string
	NotifierOpt   map[NotifierType]bool
	NotifierTypes map[NotifierType]func(Notifier) Notifier
)

var (
	// NotifierType
	NotifierTypeSlack NotifierType = "slack"
	NotifierTypeSms   NotifierType = "sms"
	NotifierTypeWa    NotifierType = "wa"

	// Notifier instantiation
	notifiers NotifierTypes = NotifierTypes{
		NotifierTypeSlack: func(n Notifier) Notifier {
			return NewSlackNotifier(n)
		},
		NotifierTypeSms: func(n Notifier) Notifier {
			return NewSmsNotifier(n)
		},
		NotifierTypeWa: func(n Notifier) Notifier {
			return NewWhatsAppNotifier(n)
		},
	}
)

func NewN(opt NotifierOpt) Notifier {
	var n Notifier = &notifier{}

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
