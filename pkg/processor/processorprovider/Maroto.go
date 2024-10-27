package processorprovider

import (
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Maroto struct {
	maroto *core.Maroto
}

func NewMaroto() *Maroto {
	// m := maroto.New()
	return nil
}

func (m *Maroto) CreateMatrixCode(codeValue string, codeProps ...*propsmapper.Rect) PDFComponent {
	cProps := propsmapper.Rect{}
	if len(codeProps) > 0 {
		cProps = *codeProps[0]
	}

	return code.NewMatrix(codeValue, props.Rect{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		JustReferenceWidth: cProps.JustReferenceWidth, Center: cProps.Center,
	})
}

func (m *Maroto) CreateBarCode(codeValue string, codeProps ...*propsmapper.Barcode) PDFComponent {
	cProps := propsmapper.Barcode{}
	if len(codeProps) > 0 {
		cProps = *codeProps[0]
	}

	return code.NewBar(codeValue, props.Barcode{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		Proportion: props.Proportion(cProps.Proportion), Center: cProps.Center, Type: barcode.Type(cProps.Type),
	})
}
