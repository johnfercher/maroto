package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type checkbox struct {
	pdf          gofpdfwrapper.Fpdf
	text         *text
	defaultColor *props.Color
}

// NewCheckbox creates a new checkbox drawer.
func NewCheckbox(pdf gofpdfwrapper.Fpdf, txt *text) *checkbox {
	return &checkbox{
		pdf:          pdf,
		text:         txt,
		defaultColor: &props.BlackColor,
	}
}

// Add draws a checkbox with label in the PDF.
func (c *checkbox) Add(label string, cell *entity.Cell, prop *props.Checkbox) {
	// Get margins to calculate absolute positions
	left, top, _, _ := c.pdf.GetMargins()

	// Calculate checkbox box position (left side of cell with padding)
	boxX := left + cell.X + prop.Left
	boxY := top + cell.Y + prop.Top
	boxSize := prop.BoxSize

	// Set draw color for the checkbox
	if prop.Color != nil {
		c.pdf.SetDrawColor(prop.Color.Red, prop.Color.Green, prop.Color.Blue)
	} else {
		c.pdf.SetDrawColor(c.defaultColor.Red, c.defaultColor.Green, c.defaultColor.Blue)
	}

	// Draw checkbox box (empty square)
	c.pdf.SetLineWidth(0.3)                       // Thin line for checkbox border
	c.pdf.Rect(boxX, boxY, boxSize, boxSize, "D") // "D" for draw (outline only)

	// If checked, draw checkmark
	if prop.Checked {
		// Draw checkmark as two lines forming a "✓" shape
		c.pdf.SetLineWidth(0.5) // Slightly thicker for checkmark

		// Start coordinates for checkmark (relative to box)
		padding := boxSize * 0.2 // 20% padding inside the box

		// First line: bottom-left to middle
		x1 := boxX + padding
		y1 := boxY + boxSize/2
		x2 := boxX + boxSize*0.4
		y2 := boxY + boxSize - padding
		c.pdf.Line(x1, y1, x2, y2)

		// Second line: middle to top-right
		x3 := x2
		y3 := y2
		x4 := boxX + boxSize - padding
		y4 := boxY + padding
		c.pdf.Line(x3, y3, x4, y4)
	}

	// Reset draw color
	c.pdf.SetDrawColor(c.defaultColor.Red, c.defaultColor.Green, c.defaultColor.Blue)
	c.pdf.SetLineWidth(0.2) // Reset to default line width

	// Calculate text position (to the right of the checkbox with spacing)
	spacing := 2.0                           // Space between checkbox and label
	textX := boxX + boxSize + spacing - left // Relative to cell

	// Create a modified cell for the text
	textCell := &entity.Cell{
		X:      textX,
		Y:      cell.Y,
		Width:  cell.Width - textX,
		Height: cell.Height,
	}

	// Create text properties from checkbox properties
	textProp := &props.Text{
		Family: prop.Family,
		Style:  prop.Style,
		Size:   prop.Size,
		Color:  prop.Color,
		Top:    prop.Top,
		Left:   0, // Already adjusted in textCell.X
		Right:  prop.Right,
		Bottom: prop.Bottom,
	}

	// Render the label
	c.text.Add(label, textCell, textProp)
}
