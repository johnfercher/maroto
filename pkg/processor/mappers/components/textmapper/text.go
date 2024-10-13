package textmapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Text struct{}

func NewText(code interface{}) (*Text, error) {
	return nil, nil
}

func (b *Text) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
