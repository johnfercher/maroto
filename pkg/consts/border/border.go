// Package border contains all border types.
package border

// Type represents a border type.
type Type int

// None is the default border type.
const None Type = 0

const (
	// Left is a border type that borders the left side.
	Left Type = 1 << iota
	// Top is a border type that borders the top side.
	Top
	// Right is a border type that borders the right side.
	Right
	// Bottom is a border type that borders the bottom side.
	Bottom
	// Full is a border type that borders all sides.
	Full = Left | Top | Right | Bottom
)

// IsValid checks if the border type is valid.
func (t Type) IsValid() bool {
	return t > None && t <= Full
}

// HasLeft checks if the border type includes left border.
func (t Type) HasLeft() bool {
	return t&Left != 0
}

// HasTop checks if the border type includes top border.
func (t Type) HasTop() bool {
	return t&Top != 0
}

// HasRight checks if the border type includes right border.
func (t Type) HasRight() bool {
	return t&Right != 0
}

// HasBottom checks if the border type includes bottom border.
func (t Type) HasBottom() bool {
	return t&Bottom != 0
}

// String returns the string representation of the border type.
func (t Type) String() string {
	if t == None {
		return ""
	}
	if t == Full {
		return "1"
	}

	result := ""
	if t.HasLeft() {
		result += "L"
	}
	if t.HasTop() {
		result += "T"
	}
	if t.HasRight() {
		result += "R"
	}
	if t.HasBottom() {
		result += "B"
	}

	return result
}
