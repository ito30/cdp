package text

type BoldText struct {
	text Text
}

func (t *BoldText) getText() string {
	return "<b>" + t.text.getText() + "</b>"
}
