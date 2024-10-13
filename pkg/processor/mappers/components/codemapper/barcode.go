// codemapper code implements creation of Barcode, MatrixCode and QrCode.
package codemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Barcode struct{}

func NewBarcode(code interface{}) (*Barcode, error) {
	return nil, nil
}

func (b *Barcode) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
