package maroto

import "github.com/jung-kurt/gofpdf"

type Page struct {
}

func (p Page) draw(fpdf *gofpdf.Fpdf) {

}

func (p Page) Header(headerClosure func(Header)) {

}

func (p Page) Footer(footerClosure func(Footer)) {

}

func (p Page) Row(rowClosure func(Row)) {

}
