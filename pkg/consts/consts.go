package consts

// Family is a representation of a family Font
type Family string

const (
	// Arial represents an arial Font
	Arial Family = "arial"
	// Helvetica represents a helvetica Font
	Helvetica Family = "helvetica"
	// Symbol represents a symbol Font
	Symbol Family = "symbol"
	// ZapBats represents a zapfdingbats Font
	ZapBats Family = "zapfdingbats"
	// Courier represents a courier Font
	Courier Family = "courier"
)

// Align is arRepresentation of a column align
type Align string

const (
	// Left represents a left horizontal align
	Left Align = "L"
	// Right represents a right horizontal align
	Right Align = "R"
	// Center represents a center horizontal and/or vertical align
	Center Align = "C"
	// Top represents a top vertical align
	Top Align = "T"
	// Bottom represents a bottom vertical align
	Bottom Align = "B"
	// Middle represents a middle align (from gofpdf)
	Middle Align = "M"
)

// Orientation is a representation of a page orientation
type Orientation string

const (
	// Portrait represents the portrait orientation.
	Portrait Orientation = "P"
	// Landscape represents the landscape orientation.
	Landscape Orientation = "L"
)

// PageSize is a representation of an page size
type PageSize string

const (
	// A3 represents DIN/ISO A3 page size
	A3 PageSize = "A3"
	// A4 represents DIN/ISO A4 page size
	A4 PageSize = "A4"
	// A5 represents DIN/ISO A5 page size
	A5 PageSize = "A5"
	// Letter represents US Letter page size
	Letter PageSize = "Letter"
	// Legal represents US Legal page size
	Legal PageSize = "Legal"
)

// Style is a representation of a style Font
type Style string

const (
	// Normal represents a normal style
	Normal Style = ""
	// Bold represents a bold style
	Bold Style = "B"
	// Italic represents a italic style
	Italic Style = "I"
	// BoldItalic represents a bold and italic style
	BoldItalic Style = "BI"
)

// Extension is a representation of a Image extension
type Extension string

const (
	// Jpg represents a jpg extension
	Jpg Extension = "jpg"
	// Png represents a png extension
	Png Extension = "png"
)
