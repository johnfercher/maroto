// Package border contains all border types.
package border

// Type represents a border type.
type Type string

const (
	// None is the default border type.
	None Type = ""
	// Full is a border type that borders all sides.
	Full Type = "1"
	// Left is a border type that borders the left side.
	Left Type = "L"
	// Top is a border type that borders the top side.
	Top Type = "T"
	// Right is a border type that borders the right side.
	Right Type = "R"
	// Bottom is a border type that borders the bottom side.
	Bottom Type = "B"
)

// IsValid checks if the border type is valid.
func (t Type) IsValid() bool {
	return t == Full || t == Left || t == Top || t == Right || t == Bottom
}
