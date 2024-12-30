package processorprovider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	processorcore "github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/loader"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Maroto struct {
	maroto     core.Maroto
	repository processorcore.ProcessorRepository
}

func NewMaroto(repository processorcore.ProcessorRepository, builder ...buildermapper.Builder) (ProcessorProvider, error) {
	cfg := config.NewBuilder()

	if len(builder) > 0 {
		var err error
		cfg, err = NewMarotoBuilder(repository, config.NewBuilder()).CreateMarotoBuilder(&builder[0])
		if err != nil {
			return nil, err
		}
	}
	m := maroto.New(cfg.Build())
	return &Maroto{maroto: m, repository: repository}, nil
}

func (m *Maroto) Generate() (core.Document, error) {
	return m.maroto.Generate()
}

func (m *Maroto) GetStructure() *node.Node[core.Structure] {
	return m.maroto.GetStructure()
}

func (m *Maroto) AddPages(pages ...ProviderComponent) (ProcessorProvider, error) {
	if len(pages) == 0 {
		return m, nil
	}
	newPages, err := convertComponentType[core.Page](pages...)
	if err != nil {
		return nil, err
	}

	m.maroto.AddPages(newPages...)
	return m, nil
}

func (m *Maroto) AddFooter(footer ...ProviderComponent) (ProcessorProvider, error) {
	if len(footer) == 0 {
		return m, nil
	}
	newFooter, err := convertComponentType[core.Row](footer...)
	if err != nil {
		return nil, err
	}

	err = m.maroto.RegisterFooter(newFooter...)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Maroto) AddHeader(header ...ProviderComponent) (ProcessorProvider, error) {
	if len(header) == 0 {
		return m, nil
	}
	newHeader, err := convertComponentType[core.Row](header...)
	if err != nil {
		return nil, err
	}

	err = m.maroto.RegisterHeader(newHeader...)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Maroto) CreatePage(components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Row](components...)
	if err != nil {
		return nil, err
	}

	return page.New().Add(newComponents...), nil
}

func (m *Maroto) CreateRow(height float64, props *propsmapper.Cell, components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Col](components...)
	if err != nil {
		return nil, err
	}

	var createdRow core.Row
	if height > 0 {
		createdRow = row.New(height).Add(newComponents...)
	} else {
		createdRow = row.New().Add(newComponents...)
	}

	if props != nil {
		return createdRow.WithStyle(createPropsCell(*props)), nil
	} else {
		return createdRow, nil
	}
}

func (m *Maroto) CreateCol(size int, props *propsmapper.Cell, components ...ProviderComponent) (ProviderComponent, error) {
	newComponents, err := convertComponentType[core.Component](components...)
	if err != nil {
		return nil, err
	}

	var createdCol core.Col
	if size > 0 {
		createdCol = col.New(size).Add(newComponents...)
	} else {
		createdCol = col.New().Add(newComponents...)
	}

	if props != nil {
		return createdCol.WithStyle(createPropsCell(*props)), nil
	} else {
		return createdCol, nil
	}
}

func (m *Maroto) CreateText(value string, textsProps ...*propsmapper.Text) ProviderComponent {
	tProps := propsmapper.Text{}
	if len(textsProps) > 0 {
		tProps = *textsProps[0]
	}

	return text.New(value, props.Text{
		Top: tProps.Top, Left: tProps.Left, Right: tProps.Right, Family: tProps.Family, Style: fontstyle.Type(tProps.Style),
		Size: tProps.Size, Align: align.Type(tProps.Align), BreakLineStrategy: breakline.Strategy(tProps.BreakLineStrategy),
		VerticalPadding: tProps.VerticalPadding, Color: (*props.Color)(tProps.Color), Hyperlink: tProps.Hyperlink,
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

func (m *Maroto) CreateImageWithLocalePath(uri *url.URL, props ...props.Rect) (ProviderComponent, error) {
	newPath := strings.TrimPrefix(uri.String(), "file://")
	return image.NewFromFile(newPath, props...), nil
}

func (m *Maroto) CreateImageWithExternalPath(path *url.URL, props ...props.Rect) (ProviderComponent, error) {
	ext, img, err := m.repository.GetDocument(path.String())
	if err != nil {
		return nil, err
	}
	return image.NewFromBytes(img, extension.Type(ext), props...), nil
}

func (m *Maroto) CreateImage(path string, propsMapperArr ...*propsmapper.Rect) (ProviderComponent, error) {
	props := createPropsRect(propsMapperArr...)
	uri, err := loader.GetResourceSource(path)
	if err != nil {
		return nil, err
	}

	if uri.Scheme == "file" {
		return m.CreateImageWithLocalePath(uri, props)
	} else {
		return m.CreateImageWithExternalPath(uri, props)
	}
}

func (m *Maroto) CreateMatrixCode(codeValue string, codeProps ...*propsmapper.Rect) ProviderComponent {
	props := createPropsRect(codeProps...)
	return code.NewMatrix(codeValue, props)
}

func (m *Maroto) CreateQrCode(codeValue string, codeProps ...*propsmapper.Rect) ProviderComponent {
	props := createPropsRect(codeProps...)
	return code.NewQr(codeValue, props)
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

func createPropsRect(propsMapperArr ...*propsmapper.Rect) props.Rect {
	propsRect := props.Rect{}
	if len(propsMapperArr) > 0 {
		propsMapper := *propsMapperArr[0]
		propsRect = props.Rect{
			Left: propsMapper.Left, Top: propsMapper.Top, Percent: propsMapper.Percent,
			JustReferenceWidth: propsMapper.JustReferenceWidth, Center: propsMapper.Center,
		}
	}
	return propsRect
}

func createPropsCell(propsCell propsmapper.Cell) *props.Cell {
	return &props.Cell{
		BackgroundColor: (*props.Color)(propsCell.BackgroundColor),
		BorderColor:     (*props.Color)(propsCell.BorderColor),
		BorderType:      border.Type(propsCell.BorderType),
		BorderThickness: propsCell.BorderThickness,
		LineStyle:       linestyle.Type(propsCell.LineStyle),
	}
}
