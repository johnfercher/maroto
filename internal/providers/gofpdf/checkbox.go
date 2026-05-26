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
		// Draw a checkmark (✓) inside the box using two connected line
		// segments: one from the left-middle area down to a bottom pivot,
		// and another from the pivot up to the top-right area.
		startX := x + prop.Size*0.2
		startY := y + prop.Size*0.5
		pivotX := x + prop.Size*0.45
		pivotY := y + prop.Size*0.75
		endX := x + prop.Size*0.8
		endY := y + prop.Size*0.2

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
