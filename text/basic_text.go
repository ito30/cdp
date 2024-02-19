package text

type BasicText struct {
	text string
}

func (b *BasicText) getText() string {
	return b.text
}
