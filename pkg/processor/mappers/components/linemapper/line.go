// Package line implements creation of lines.
package linemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Line struct{}

func NewLine(code interface{}) (*Line, error) {
	return nil, nil
}

func (b *Line) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
