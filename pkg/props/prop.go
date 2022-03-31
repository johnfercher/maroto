package props

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
)

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
	// Top is space between the upper cell limit to the barcode, if align is not center.
	Top float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style consts.Style
	// Size of the text.
	Size float64
	// Align of the text.
	Align consts.Align
	// Extrapolate define if the text will automatically add a new line when.
	// text reach the right cell boundary.
	Extrapolate bool
	// VerticalPadding define an additional space between lines.
	VerticalPadding float64
	// Color define the font color.
	Color color.Color
}

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color color.Color
	// Style define the line style (solid, dashed or dotted).
	Style consts.LineStyle
	// Width define the line width (it cannot be greater than line height)
	Width float64
}

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style consts.Style
	// Size of the text.
	Size float64
	// Color define the font color.
	Color color.Color
}

// TableListContent represents properties from a line (header/content) from a TableList.
type TableListContent struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style consts.Style
	// Size of the text.
	Size float64
	// Color define the font color.
	Color color.Color
	// GridSizes is the custom properties of the size of the grid
	// the sum of the values cannot be greater than 12, if this
	// value is not provided the width of all columns will be the
	// same.
	GridSizes []uint
}

// TableList represents properties from a TableList.
type TableList struct {
	// HeaderProp is the custom properties of the text inside
	// the headers.
	HeaderProp TableListContent
	// ContentProp is the custom properties of the text inside
	// the contents.
	ContentProp TableListContent
	// Align is the align of the text (header and content) inside the columns.
	Align consts.Align
	// AlternatedBackground define the background color from even rows
	// i.e rows with index (0, 2, 4, ..., N) will have background colorized,
	// rows with index (1, 3, 5, ..., N) will not.
	AlternatedBackground *color.Color
	// HeaderContentSpace is the space between the header and the contents.
	HeaderContentSpace float64
	// Line adds a line after every content-row to separate rows. The line's spaceHeight is set to 1.0.
	Line bool
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
func (s *Text) MakeValid(defaultFamily string) {
	minValue := 0.0
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = defaultFamily
	}

	if s.Style == "" {
		s.Style = consts.Normal
	}

	if s.Align == "" {
		s.Align = consts.Left
	}

	if s.Size == undefinedValue {
		s.Size = 10.0
	}

	if s.Top < minValue {
		s.Top = minValue
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
		s.Style = consts.Bold
	}

	if s.Size == undefinedValue {
		s.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font.
func (s *Font) ToTextProp(align consts.Align, top float64, extrapolate bool, verticalPadding float64) Text {
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

	textProp.MakeValid(s.Family)

	return textProp
}

// ToTextProp from Font return a TableListContent based on Font.
func (s *TableListContent) ToTextProp(align consts.Align, top float64, extrapolate bool, verticalPadding float64) Text {
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

	textProp.MakeValid(s.Family)

	return textProp
}

// MakeValid from TableList define default values for a TableList.
func (s *TableList) MakeValid(header []string, defaultFamily string) {
	undefinedValue := 0.0
	if s.HeaderProp.Size == undefinedValue {
		s.HeaderProp.Size = 10.0
	}

	if s.HeaderProp.Family == "" {
		s.HeaderProp.Family = defaultFamily
	}

	if s.HeaderProp.Style == "" {
		s.HeaderProp.Style = consts.Bold
	}

	if len(s.HeaderProp.GridSizes) == 0 {
		gridSize := uint(consts.MaxGridSum / float64(len(header)))
		s.HeaderProp.GridSizes = []uint{}

		for range header {
			s.HeaderProp.GridSizes = append(s.HeaderProp.GridSizes, gridSize)
		}
	}

	if s.Align == "" {
		s.Align = consts.Left
	}

	if s.ContentProp.Size == undefinedValue {
		s.ContentProp.Size = 10.0
	}

	if s.ContentProp.Family == "" {
		s.ContentProp.Family = defaultFamily
	}

	if s.ContentProp.Style == "" {
		s.ContentProp.Style = consts.Normal
	}

	if len(s.ContentProp.GridSizes) == 0 {
		gridSize := uint(consts.MaxGridSum / float64(len(header)))
		s.ContentProp.GridSizes = []uint{}

		for range header {
			s.ContentProp.GridSizes = append(s.ContentProp.GridSizes, gridSize)
		}
	}

	if s.HeaderContentSpace == undefinedValue {
		s.HeaderContentSpace = 4.0
	}
}

// MakeValid from Line define default values for a Line.
func (s *Line) MakeValid(spaceHeight float64) {
	if s.Style == "" {
		s.Style = consts.Solid
	}

	if s.Width == 0 {
		s.Width = consts.DefaultLineWidth
	}

	if s.Width > spaceHeight {
		s.Width = spaceHeight
	}
}
