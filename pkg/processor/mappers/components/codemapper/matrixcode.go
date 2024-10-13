// Package codemapper implements creation of Barcode, MatrixCode and QrCode.
package codemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Matrixcode struct{}

func NewMatrixcode(code interface{}) (*Matrixcode, error) {
	return nil, nil
}

func (b *Matrixcode) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
