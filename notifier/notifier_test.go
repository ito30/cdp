package notifier

import (
	"fmt"
	"testing"
)

func TestNotifier(t *testing.T) {
	message := "Hi you got notification from John Doe!"
	notifier := &notifier{}

	tests := map[string]struct {
		notifier func() INotifier
	}{
		"slack": {
			notifier: func() INotifier {
				return NewSlackNotifier(notifier)
			},
		},
		"slack and sms": {
			notifier: func() INotifier {
				slackNotifier := NewSlackNotifier(notifier)
				return NewSmsNotifier(slackNotifier)
			},
		},
		"slack and sms and wa": {
			notifier: func() INotifier {
				slackNotifier := NewSlackNotifier(notifier)
				slackAndSmsNotifier := NewSmsNotifier(slackNotifier)
				return NewWhatsAppNotifier(slackAndSmsNotifier)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.notifier().Send(message)
			fmt.Println()
		})
	}
}
