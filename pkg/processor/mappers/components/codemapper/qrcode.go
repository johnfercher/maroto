// Package codemapper implements creation of Barcode, MatrixCode and QrCode.
package codemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Qrcode struct{}

func NewQrcode(code interface{}) (*Qrcode, error) {
	return nil, nil
}

func (b *Qrcode) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
