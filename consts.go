package maroto

type Family int

const (
	Arial     Family = 0
	Helvetica Family = 1
	Symbol    Family = 2
	ZapBats   Family = 3
	Courier   Family = 4
)

func GetFamilyString(font Family) string {
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

type Align int

const (
	Left   Align = 0
	Right  Align = 1
	Center Align = 2
	Top    Align = 3
	Bottom Align = 4
)

type Orientation int

const (
	Vertical Orientation = 0
)

type PageSize int

const (
	A4 PageSize = 0
)

type Style int

const (
	Normal     Style = 0
	Bold       Style = 1
	Italic     Style = 2
	BoldItalic Style = 3
)

func GetStyleString(style Style) string {
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

type Extension int

const (
	Jpg Extension = 0
	Png Extension = 1
)

func GetExtensionString(extension Extension) string {
	switch extension {
	case Png:
		return "png"
	default:
		return "jpg"
	}
}
