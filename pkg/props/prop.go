package props

import (
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

type Style struct {
	BackgroundColor *color.Color
	BorderColor     *color.Color
	Border          border.Type
}

// Proportion represents a proportion from a rectangle, example: 16x9, 4x3...
type Proportion struct {
	// Width from the rectangle: Barcode, image and etc.
	Width float64
	// Height from the rectangle: Barcode, image and etc.
	Height float64
}

// Barcode represents properties from a barcode inside a cell.
type Barcode struct {
	// Left is the space between the left cell boundary to the barcode, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the barcode will occupy the cell,
	// ex 100%: The barcode will fulfill the entire cell
	// ex 50%: The greater side from the barcode will have half the size of the cell.
	Percent float64
	// Proportion is the proportion between size of the barcode.
	// Ex: 16x9, 4x3...
	Proportion Proportion
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}

// Rect represents properties from a rectangle (Image, QrCode or Barcode) inside a cell.
type Rect struct {
	// Left is the space between the left cell boundary to the rectangle, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the rectangle will occupy the cell,
	// ex 100%: The rectangle will fulfill the entire cell
	// ex 50%: The greater side from the rectangle will have half the size of the cell.
	Percent float64
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Left is the minimal amount of space between the left cell boundary and the text.
	Left float64
	// Right is the minimal amount of space between the right cell boundary and the text.
	Right float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Align of the text.
	Align align.Type
	// Extrapolate define if the text will automatically add a new line when.
	// text reach the right cell boundary.
	Extrapolate bool
	// VerticalPadding define an additional space between lines.
	VerticalPadding float64
	// Color define the fontstyle color.
	Color *color.Color
}

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color color.Color
	// Style define the line style (solid, dashed or dotted).
	Style linestyle.Type
	// Width define the line width (it cannot be greater than line height)
	Width float64
}

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the fontstyle color.
	Color *color.Color
}

// MakeValid from Rect will make the properties from a rectangle reliable to fit inside a cell
// and define default values for a rectangle.
func (s *Rect) MakeValid() {
	minPercentage := 0.0
	maxPercentage := 100.0
	minValue := 0.0

	if s.Percent <= minPercentage || s.Percent > maxPercentage {
		s.Percent = maxPercentage
	}

	if s.Center {
		s.Left = 0
		s.Top = 0
	}

	if s.Left < minValue {
		s.Left = minValue
	}

	if s.Top < minValue {
		s.Top = minValue
	}
}

// MakeValid from Barcode will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode.
func (s *Barcode) MakeValid() {
	minPercentage := 0.0
	maxPercentage := 100.0
	minValue := 0.0

	if s.Percent <= minPercentage || s.Percent > maxPercentage {
		s.Percent = maxPercentage
	}

	if s.Center {
		s.Left = 0
		s.Top = 0
	}

	if s.Left < minValue {
		s.Left = minValue
	}

	if s.Top < minValue {
		s.Top = minValue
	}

	if s.Proportion.Width <= 0 {
		s.Proportion.Width = 1
	}

	if s.Proportion.Height <= 0 {
		s.Proportion.Height = 1
	}

	maxHeightProportionBasedOnWidth := 0.20
	minHeightProportionBasedOnWidth := 0.10

	if s.Proportion.Height > s.Proportion.Width*maxHeightProportionBasedOnWidth {
		s.Proportion.Height = s.Proportion.Width * maxHeightProportionBasedOnWidth
	} else if s.Proportion.Height < s.Proportion.Width*minHeightProportionBasedOnWidth {
		s.Proportion.Height = s.Proportion.Width * minHeightProportionBasedOnWidth
	}
}

// MakeValid from Text define default values for a Text.
func (s *Text) MakeValid(font *Font) {
	minValue := 0.0
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = font.Family
	}

	if s.Style == "" {
		s.Style = font.Style
	}

	if s.Size == undefinedValue {
		s.Size = font.Size
	}

	if s.Color == nil {
		s.Color = font.Color
	}

	if s.Align == "" {
		s.Align = align.Center
	}

	if s.Top < minValue {
		s.Top = minValue
	}

	if s.Left < minValue {
		s.Left = minValue
	}

	if s.Right < minValue {
		s.Right = minValue
	}

	if s.VerticalPadding < 0 {
		s.VerticalPadding = 0
	}
}

// MakeValid from Font define default values for a Signature.
func (s *Font) MakeValid(defaultFamily string) {
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = defaultFamily
	}

	if s.Style == "" {
		s.Style = fontstyle.Bold
	}

	if s.Size == undefinedValue {
		s.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font.
func (s *Font) ToTextProp(align align.Type, top float64, extrapolate bool, verticalPadding float64) Text {
	textProp := Text{
		Family:          s.Family,
		Style:           s.Style,
		Size:            s.Size,
		Align:           align,
		Top:             top,
		Extrapolate:     extrapolate,
		VerticalPadding: verticalPadding,
		Color:           s.Color,
	}

	textProp.MakeValid(s)

	return textProp
}

// MakeValid from Line define default values for a Line.
func (s *Line) MakeValid(spaceHeight float64) {
	if s.Style == "" {
		s.Style = linestyle.Solid
	}

	if s.Width == 0 {
		s.Width = linestyle.DefaultLineWidth
	}

	if s.Width > spaceHeight {
		s.Width = spaceHeight
	}
}
