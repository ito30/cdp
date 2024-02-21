package notifier

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNotifier(t *testing.T) {
	tests := map[string]struct {
		userCfg  string
		notifier func(userCfg string) Notifier
	}{
		"slack": {
			userCfg: `
				{
					"slack": true
				}
			`,
			notifier: func(userCfg string) Notifier {
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
			notifier: func(userCfg string) Notifier {
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
			notifier: func(userCfg string) Notifier {
				opt := ParseToObj[NotifierOpt](userCfg)
				return NewN(opt)
			},
		},
	}

	for _, tt := range tests {
		tt.notifier(tt.userCfg).Send("Kamu dapat 3000 Gopay coins!")
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
