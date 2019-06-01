package maroto

type Family string

const (
	Arial     Family = "arial"
	Helvetica Family = "helvetica"
	Symbol    Family = "symbol"
	ZapBats   Family = "zapfdingbats"
	Courier   Family = "courier"
)

type Align int

const (
	Left   Align = 0
	Right  Align = 1
	Center Align = 2
	Top    Align = 3
	Bottom Align = 4
)

type Orientation string

const (
	Portrait  Orientation = "P"
	Landscape Orientation = "L"
)

type PageSize string

const (
	A3     PageSize = "A3"
	A4     PageSize = "A4"
	A5     PageSize = "A5"
	Letter PageSize = "Letter"
	Legal  PageSize = "Legal"
)

type Style string

const (
	Normal     Style = ""
	Bold       Style = "B"
	Italic     Style = "I"
	BoldItalic Style = "BI"
)

type Extension string

const (
	Jpg Extension = "jpg"
	Png Extension = "png"
)
