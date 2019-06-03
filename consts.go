package maroto

// Representation of a family Font
type Family string

const (
	// Represents an arial Font
	Arial Family = "arial"
	// Represents a helvetica Font
	Helvetica Family = "helvetica"
	// Represents a symbol Font
	Symbol Family = "symbol"
	// Represents a zapfdingbats Font
	ZapBats Family = "zapfdingbats"
	// Represents a courier Font
	Courier Family = "courier"
)

// Representation of a column align
type Align string

const (
	// Represents a left align
	Left Align = "L"
	// Represents a right align
	Right Align = "R"
	// Represents a center align
	Center Align = "C"
	// Represents a top align
	Top Align = "T"
	// Represents a bottom align
	Bottom Align = "B"
	// Represents a middle align
	Middle Align = "M"
)

// Representation of a page orientation
type Orientation string

const (
	// Represents the portrait orientation.
	Portrait Orientation = "P"
	// Represents the landscape orientation.
	Landscape Orientation = "L"
)

// Representation of an page size
type PageSize string

const (
	// Represents DIN/ISO A3 page size
	A3 PageSize = "A3"
	// Represents DIN/ISO A4 page size
	A4 PageSize = "A4"
	// Represents DIN/ISO A5 page size
	A5 PageSize = "A5"
	// Represents US Letter page size
	Letter PageSize = "Letter"
	// Represents US Legal page size
	Legal PageSize = "Legal"
)

// Representation of a style Font
type Style string

const (
	// Represents a normal style
	Normal Style = ""
	// Represents a bold style
	Bold Style = "B"
	// Represents a italic style
	Italic Style = "I"
	// Represents a bold and italic style
	BoldItalic Style = "BI"
)

// Representation of a Image extension
type Extension string

const (
	// Represents a jpg extension
	Jpg Extension = "jpg"
	// Represents a png extension
	Png Extension = "png"
)
