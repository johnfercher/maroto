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
		// Draw checkmark (✓) inside the box using two connected line
		// segments: a short stroke from the left side down to a low pivot
		// near the bottom-center, and a longer stroke up to the top-right.
		startX, startY := x+prop.Size*0.2, y+prop.Size*0.5
		pivotX, pivotY := x+prop.Size*0.45, y+prop.Size*0.75
		endX, endY := x+prop.Size*0.85, y+prop.Size*0.25

		c.pdf.Line(startX, startY, pivotX, pivotY)
		c.pdf.Line(pivotX, pivotY, endX, endY)
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
