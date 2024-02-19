package notifier

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
