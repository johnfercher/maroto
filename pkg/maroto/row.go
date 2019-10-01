package maroto

import "github.com/jung-kurt/gofpdf"

type Row struct {
	Height int32
}

func (r Row) draw(fpdf *gofpdf.Fpdf) {

}

func (r Row) Col(colClosure func(col Col)) {

}

type Header struct {
	Row
	Repeat bool
}

type Footer struct {
	Header
	Absolute bool
}
