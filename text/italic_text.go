package text

type ItalicText struct {
	text Text
}

func (t *ItalicText) getText() string {
	return "<i>" + t.text.getText() + "</i>"
}
