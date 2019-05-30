package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

type Font interface {
	SetFamily(family Family)
	SetStyle(style Style)
	SetSize(size float64)
	SetFont(family Family, style Style, size float64)
	GetFamily() Family
	GetStyle() Style
	GetSize() float64
	GetFont() (Family, Style, float64)
	GetStyleString(style Style) string
	GetFamilyString(font Family) string
}

type font struct {
	pdf    gofpdf.Pdf
	size   float64
	family Family
	style  Style
}

func NewFont(pdf gofpdf.Pdf, size float64, family Family, style Style) Font {
	return &font{
		pdf,
		size,
		family,
		style,
	}
}

func (f *font) GetFamily() Family {
	return f.family
}

func (f *font) GetStyle() Style {
	return f.style
}

func (f *font) GetSize() float64 {
	return f.size
}

func (f *font) GetFont() (Family, Style, float64) {
	return f.family, f.style, f.size
}

func (f *font) SetFamily(family Family) {
	f.family = family
	familyString := f.GetFamilyString(f.family)
	styleString := f.GetStyleString(f.style)

	f.pdf.SetFont(familyString, styleString, f.size)
}

func (f *font) SetStyle(style Style) {
	f.style = style
	styleString := f.GetStyleString(f.style)

	f.pdf.SetFontStyle(styleString)
}

func (f *font) SetSize(size float64) {
	f.size = size
	f.pdf.SetFontSize(f.size)
}

func (f *font) SetFont(family Family, style Style, size float64) {
	f.family = family
	f.style = style
	f.size = size

	familyString := f.GetFamilyString(f.family)
	styleString := f.GetStyleString(f.style)

	f.pdf.SetFont(familyString, styleString, f.size)
}

func (f *font) GetFamilyString(font Family) string {
	switch font {
	case Courier:
		return "courier"
	case Helvetica:
		return "helvetica"
	case Symbol:
		return "symbol"
	case ZapBats:
		return "zapfdingbats"
	default:
		return "arial"
	}
}

func (f *font) GetStyleString(style Style) string {
	switch style {
	case Italic:
		return "I"
	case Bold:
		return "B"
	case BoldItalic:
		return "BI"
	default:
		return ""
	}
}
