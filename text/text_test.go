package text

import (
	"fmt"
	"testing"
)

func TestNotifier(t *testing.T) {
	tests := map[string]struct {
		text func() Text
	}{
		"basic text": {
			text: func() Text {
				return &BasicText{text: "Hello, world!"}
			},
		},
		"bold text": {
			text: func() Text {
				basic := &BasicText{text: "Hello, world!"}
				return &BoldText{text: basic}
			},
		},
		"bold & italic text": {
			text: func() Text {
				basic := &BasicText{text: "Hello, world!"}
				basicAndBold := &BoldText{text: basic}

				return &ItalicText{
					text: basicAndBold,
				}
			},
		},
		"bold & italic & underline text": {
			text: func() Text {
				basic := &BasicText{
					text: "Hello, world!",
				}
				basicAndBold := &BoldText{
					text: basic,
				}

				basicAndBoldAndUnderline := &ItalicText{
					text: basicAndBold,
				}

				return &UnderlineText{
					text: basicAndBoldAndUnderline,
				}
			},
		},
	}

	for _, tt := range tests {
		text := tt.text().getText()
		fmt.Printf("Output text: %s\n", text)
	}
}
