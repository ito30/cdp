package text

type UnderlineText struct {
	text Text
}

func (t *UnderlineText) getText() string {
	return "<u>" + t.text.getText() + "</u>"
}
