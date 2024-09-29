// barcode is the package responsible for mapping barcode settings
package barcode

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/barcode"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type BarCode struct {
	Props     propsmapper.BarCodeProps `json:"props"`
	SourceKey string                   `json:"source_key"`
}

// generate is responsible for the builder barcode according to the submitted content
func (b *BarCode) Generate(content map[string]interface{}) (components.Component, error) {
	value, ok := content[b.SourceKey]
	if !ok {
		return nil, fmt.Errorf("barcode model needs source key %s, but no content with that key was found", b.SourceKey)
	}

	code, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("resource %s does not have a valid value for the barcode component", b.SourceKey)
	}

	return barcode.NewBarCode(props.BarCodeProps{Align: b.Props.Align}, code), nil
}
