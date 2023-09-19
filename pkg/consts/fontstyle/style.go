package fontstyle

// Type is a representation of a style Font.
type Type string

const (
	// Normal represents a normal style.
	Normal Type = ""
	// Bold represents a bold style.
	Bold Type = "B"
	// Italic represents a italic style.
	Italic Type = "I"
	// BoldItalic represents a bold and italic style.
	BoldItalic Type = "BI"
)

func (s Type) IsValid() bool {
	return s == Normal || s == Italic || s == BoldItalic || s == Bold
}
