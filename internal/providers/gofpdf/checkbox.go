package gofpdf

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const labelGap = 1.0

// Check mark geometry expressed as ratios of the checkbox size. The mark
// is composed of two connected line segments forming a "✓":
//   - start (left)  → mid (bottom of the V)
//   - mid           → end (upper-right tip)
//
// The ratios use eighths so multiplying by any size produces an exact
// float64 result (avoids floating-point drift in tests).
const (
	checkStartXRatio = 0.25  // 2/8
	checkStartYRatio = 0.625 // 5/8
	checkMidXRatio   = 0.375 // 3/8
	checkMidYRatio   = 0.75  // 6/8
	checkEndXRatio   = 0.875 // 7/8
	checkEndYRatio   = 0.25  // 2/8
)

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
		// Draw a check mark (✓) inside the box using two connected line
		// segments: a short stroke down-right to the mid point, then a
		// longer stroke up-right to the top-right area of the box.
		startX := x + prop.Size*checkStartXRatio
		startY := y + prop.Size*checkStartYRatio
		midX := x + prop.Size*checkMidXRatio
		midY := y + prop.Size*checkMidYRatio
		endX := x + prop.Size*checkEndXRatio
		endY := y + prop.Size*checkEndYRatio

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
