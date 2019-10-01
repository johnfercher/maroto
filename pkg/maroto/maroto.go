package maroto

import "github.com/jung-kurt/gofpdf"

type Maroto struct {
	Layout Layout
	fpdf   *gofpdf.Fpdf
}

func (m *Maroto) Init() {
	l := m.Layout
	m.fpdf = gofpdf.New(string(l.Orientation), "mm", string(l.PageSize), "")
	m.fpdf.SetMargins(l.Margins.Left, l.Margins.Right, l.Margins.Top)
}

func (m *Maroto) Page(pageClosure func(Page)) {

	if m.fpdf == nil {
		panic("Maroto was not initialized!")
	}

	page := Page{}
	pageClosure(page)
	page.draw(m.fpdf)
}
