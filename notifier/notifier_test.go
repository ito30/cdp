package notifier

import (
	"encoding/json"
	"fmt"
	"testing"
)

// type NotifierOpt struct {
// 	Slack bool `json:"slack"`
// 	Sms   bool `json:"sms"`
// 	Wa    bool `json:"wa"`
// }

type (
	NotifierType string
	NotifierOpt  map[NotifierType]bool
)

var (
	// NotifierType
	NotifierTypeSlack NotifierType = "slack"
	NotifierTypeSms   NotifierType = "sms"
	NotifierTypeWa    NotifierType = "wa"

	// Notifier instantiation
	notifiers map[NotifierType]func(INotifier) INotifier = map[NotifierType]func(INotifier) INotifier{
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
		if nt, ok := notifiers[nType]; ok && active {
			n = nt(n)
		}
	}

	return n
}

func TestNotifier(t *testing.T) {
	message := "Hi you got notification from John Doe!"

	tests := map[string]struct {
		userCfg  string
		notifier func(userCfg string) INotifier
	}{
		"slack": {
			userCfg: `
				{
					"slack": true
				}
			`,
			notifier: func(userCfg string) INotifier {
				opt := ParseToObj[NotifierOpt](userCfg)
				return NewN(opt)
			},
		},
		"slack ~ sms": {
			userCfg: `
				{
					"slack": true,
					"sms": true
				}
			`,
			notifier: func(userCfg string) INotifier {
				opt := ParseToObj[NotifierOpt](userCfg)
				return NewN(opt)
			},
		},
		"slack ~ sms ~ wa": {
			userCfg: `
				{
					"slack": true,
					"sms": true,
					"wa": true
				}
			`,
			notifier: func(userCfg string) INotifier {
				opt := ParseToObj[NotifierOpt](userCfg)
				return NewN(opt)
			},
		},
	}

	for _, tt := range tests {
		tt.notifier(tt.userCfg).Send(message)
		fmt.Println()
	}
}

func ParseToObj[T any](jsonStr string) T {
	var obj T

	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		panic(err)
	}

	return obj
}
