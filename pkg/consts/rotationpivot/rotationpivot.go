// Package rotationpivot defines the anchor points used when rotating text.
package rotationpivot

// Type selects the horizontal anchor of the text glyph block during rotation.
type Type string

const (
	// Start anchors rotation at the leading edge of the text (left side for
	// left-to-right scripts).
	Start Type = "start"
	// Center anchors rotation at the horizontal center of the text. Default.
	Center Type = "center"
	// End anchors rotation at the trailing edge of the text (right side for
	// left-to-right scripts).
	End Type = "end"
)

// VerticalType selects the vertical anchor of the text glyph block during
// rotation. For multi-line text the anchor refers to the whole block, not a
// single line.
type VerticalType string

const (
	// Top anchors rotation at the top of the text block.
	Top VerticalType = "top"
	// Middle anchors rotation at the vertical center of the text block. Default.
	Middle VerticalType = "middle"
	// Bottom anchors rotation at the bottom of the text block.
	Bottom VerticalType = "bottom"
)

// Pivot is the combined horizontal + vertical anchor used when rotating text.
type Pivot struct {
	Horizontal Type
	Vertical   VerticalType
}
