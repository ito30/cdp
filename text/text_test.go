package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifier(t *testing.T) {
	tests := map[string]struct {
		text func() Text
		want string
	}{
		"basic text": {
			text: func() Text {
				return &BasicText{text: "Hello, world!"}
			},
			want: "Hello, world!",
		},
		"bold text": {
			text: func() Text {
				basic := &BasicText{text: "Hello, world!"}
				return &BoldText{text: basic}
			},
			want: "<b>Hello, world!</b>",
		},
		"bold & italic text": {
			text: func() Text {
				basic := &BasicText{text: "Hello, world!"}
				basicAndBold := &BoldText{text: basic}

				return &ItalicText{
					text: basicAndBold,
				}
			},
			want: "<i><b>Hello, world!</b></i>",
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
			want: "<u><i><b>Hello, world!</b></i></u>",
		},
	}

	for _, tt := range tests {
		text := tt.text().getText()
		fmt.Printf("Output text: %s\n", text)

		assert.Equal(t, tt.want, text)
	}
}
