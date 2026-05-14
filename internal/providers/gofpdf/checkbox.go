package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const labelGap = 1.0

type Checkbox struct {
	pdf  gofpdfwrapper.Fpdf
	font core.Font
}

// NewCheckbox create a Checkbox.
func NewCheckbox(pdf gofpdfwrapper.Fpdf, font core.Font) *Checkbox {
	return &Checkbox{pdf: pdf, font: font}
}

// Add a checkbox with a label inside a cell.
func (c *Checkbox) Add(label string, cell *entity.Cell, prop *props.Checkbox) {
	left, top, _, _ := c.pdf.GetMargins()

	x := cell.X + prop.Left + left
	y := cell.Y + prop.Top + top

	// Draw the checkbox square border
	c.pdf.Rect(x, y, prop.Size, prop.Size, "D")

	if prop.Checked {
		// Draw checkmark (✓) inside the box using two lines.
		// The mark is composed of a short descending stroke followed
		// by a longer ascending stroke, joined at the lower vertex.
		// Coordinates are expressed as fractions of prop.Size so the
		// checkmark scales with the box.
		startX := x + 0.20*prop.Size
		startY := y + 0.50*prop.Size
		midX := x + 0.45*prop.Size
		midY := y + 0.75*prop.Size
		endX := x + 0.85*prop.Size
		endY := y + 0.25*prop.Size

		c.pdf.Line(startX, startY, midX, midY)
		c.pdf.Line(midX, midY, endX, endY)
	}

	// Draw label to the right of the checkbox, vertically centered
	if label != "" {
		family, style, size := c.font.GetFont()
		fontHeight := c.font.GetHeight(family, style, size)

		labelX := x + prop.Size + labelGap
		labelY := y + prop.Size/2 + fontHeight/2

		c.pdf.Text(labelX, labelY, label)
	}
}
