// Package signature implements creation of signatures.
package signaturemapper

import "github.com/johnfercher/maroto/v2/pkg/processor/components"

type Signature struct{}

func NewSignature(code interface{}) (*Signature, error) {
	return nil, nil
}

func (b *Signature) Generate(content map[string]interface{}) (components.PdfComponent, error) {
	return nil, nil
}
