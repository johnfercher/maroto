package processorprovider

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	marototext "github.com/johnfercher/maroto/v2/pkg/components/text"
	marotoprops "github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
)

type Maroto struct {
	maroto *core.Maroto
}

func NewMaroto() *Maroto {
	m := maroto.New()
	return &Maroto{maroto: &m}
}

func (m *Maroto) GeneratePdf() ([]byte, error) {
	doc, err := (*m.maroto).Generate()
	if err != nil {
		return nil, err
	}
	doc.Save("docs/assets/pdf/backgroundv2.pdf")
	return doc.GetBytes(), nil
}

func (m *Maroto) ConfigureBuilder(builder builder.Builder) ProcessorProvider {
	return nil
}

func (m *Maroto) RegisterHeader(rows ...core.Row) ProcessorProvider {
	(*m.maroto).RegisterHeader(rows...)
	return m
}

func (m *Maroto) RegisterFooter(rows ...core.Row) ProcessorProvider {
	(*m.maroto).RegisterFooter(rows...)
	return m
}

func (m *Maroto) CreatePage(components ...core.Row) core.Page {
	newPage := page.New().Add(components...)
	(*m.maroto).AddPages(newPage)
	return newPage
}

func (m *Maroto) CreateRow(cols ...core.Col) core.Row {
	return row.New().Add(cols...)
}

func (m *Maroto) CreateCol(size int, components ...core.Component) core.Col {
	return col.New(size).Add(components...)
}

func (m *Maroto) CreateText(value string, props props.TextProps) core.Component {
	return marototext.New(value, marotoprops.Text{Align: align.Type(props.Align)})
}

func (m *Maroto) CreateBarCode(codeValue string, props props.BarCodeProps) core.Component {
	return code.NewBar(codeValue)
}
