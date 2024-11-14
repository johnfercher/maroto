package processorprovider

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
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

func convertComponentType[T any](components ...ProviderComponent) ([]T, error) {
	newComponents := make([]T, len(components))
	for i, component := range components {
		validComponent, ok := component.(T)
		if !ok {
			return nil, fmt.Errorf("could not convert pdf components to a valid type")
		}
		newComponents[i] = validComponent
	}
	return newComponents, nil
}

func (m *Maroto) CreatePage(components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Row](components...)
	if err != nil {
		return nil, err
	}

	return page.New().Add(newComponents...), nil
}

func (m *Maroto) CreateRow(height float64, components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Col](components...)
	if err != nil {
		return nil, err
	}

	return row.New(height).Add(newComponents...), nil
}

func (m *Maroto) CreateCol(size int, components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Component](components...)
	if err != nil {
		return nil, err
	}

	return col.New(size).Add(newComponents...), nil
}

func (m *Maroto) CreateText(value string, textsProps ...*propsmapper.Text) ProviderComponent {
	tProps := propsmapper.Text{}
	if len(textsProps) > 0 {
		tProps = *textsProps[0]
	}

	return text.New(value, props.Text{
		Top: tProps.Top, Left: tProps.Left, Right: tProps.Right, Family: tProps.Family, Style: fontstyle.Type(tProps.Style),
		Size: tProps.Size, Align: align.Type(tProps.Align), BreakLineStrategy: breakline.Strategy(tProps.BreakLineStrategy),
		VerticalPadding: tProps.VerticalPadding, Color: (*props.Color)(tProps.Color), Hyperlink: &tProps.Hyperlink,
	})
}

func (m *Maroto) CreateSignature(value string, signaturesProps ...*propsmapper.Signature) ProviderComponent {
	sProps := propsmapper.Signature{}
	if len(signaturesProps) > 0 {
		sProps = *signaturesProps[0]
	}

	return signature.New(value, props.Signature{
		FontFamily: sProps.FontFamily, FontStyle: fontstyle.Type(sProps.FontStyle), FontSize: sProps.FontSize,
		FontColor: (*props.Color)(sProps.FontColor), LineColor: (*props.Color)(sProps.LineColor), LineStyle: sProps.LineStyle,
		LineThickness: sProps.LineThickness, SafePadding: sProps.SafePadding,
	})
}

func (m *Maroto) CreateLine(lineProps ...*propsmapper.Line) ProviderComponent {
	lProps := propsmapper.Line{}
	if len(lineProps) > 0 {
		lProps = *lineProps[0]
	}

	return line.New(props.Line{
		Color: (*props.Color)(lProps.Color), Style: linestyle.Type(lProps.Style), Thickness: lProps.Thickness,
		Orientation: orientation.Type(lProps.Orientation), OffsetPercent: lProps.OffsetPercent, SizePercent: lProps.SizePercent,
	})
}

func (m *Maroto) CreateImage(img []byte, ext string, imgProps ...*propsmapper.Rect) ProviderComponent {
	cProps := propsmapper.Rect{}
	if len(imgProps) > 0 {
		cProps = *imgProps[0]
	}

	return image.NewFromBytes(img, extension.Type(ext), props.Rect{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		JustReferenceWidth: cProps.JustReferenceWidth, Center: cProps.Center,
	})
}

func (m *Maroto) CreateMatrixCode(codeValue string, codeProps ...*propsmapper.Rect) ProviderComponent {
	cProps := propsmapper.Rect{}
	if len(codeProps) > 0 {
		cProps = *codeProps[0]
	}

	return code.NewMatrix(codeValue, props.Rect{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		JustReferenceWidth: cProps.JustReferenceWidth, Center: cProps.Center,
	})
}

func (m *Maroto) CreateQrCode(codeValue string, codeProps ...*propsmapper.Rect) ProviderComponent {
	cProps := propsmapper.Rect{}
	if len(codeProps) > 0 {
		cProps = *codeProps[0]
	}

	return code.NewQr(codeValue, props.Rect{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		JustReferenceWidth: cProps.JustReferenceWidth, Center: cProps.Center,
	})
}

func (m *Maroto) CreateBarCode(codeValue string, codeProps ...*propsmapper.Barcode) ProviderComponent {
	cProps := propsmapper.Barcode{}
	if len(codeProps) > 0 {
		cProps = *codeProps[0]
	}

	return code.NewBar(codeValue, props.Barcode{
		Left: cProps.Left, Top: cProps.Top, Percent: cProps.Percent,
		Proportion: props.Proportion(cProps.Proportion), Center: cProps.Center, Type: barcode.Type(cProps.Type),
	})
}
