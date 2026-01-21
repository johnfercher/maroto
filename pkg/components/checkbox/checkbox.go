// Package checkbox implements creation of checkboxes.
package checkbox

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Checkbox struct {
	label  string
	prop   props.Checkbox
	config *entity.Config
}

// New is responsible to create an instance of a Checkbox.
func New(label string, ps ...props.Checkbox) core.Component {
	checkboxProp := props.Checkbox{}
	if len(ps) > 0 {
		checkboxProp = ps[0]
	}

	return &Checkbox{
		label: label,
		prop:  checkboxProp,
	}
}

// NewCol is responsible to create an instance of a Checkbox wrapped in a Col.
func NewCol(size int, label string, ps ...props.Checkbox) core.Col {
	checkbox := New(label, ps...)
	return col.New(size).Add(checkbox)
}

// NewAutoRow is responsible for creating an instance of Checkbox grouped in a Row with automatic height.
func NewAutoRow(label string, ps ...props.Checkbox) core.Row {
	c := New(label, ps...)
	column := col.New().Add(c)
	return row.New().Add(column)
}

// NewRow is responsible to create an instance of a Checkbox wrapped in a Row.
func NewRow(height float64, label string, ps ...props.Checkbox) core.Row {
	c := New(label, ps...)
	column := col.New().Add(c)
	return row.New(height).Add(column)
}

// GetStructure returns the Structure of a Checkbox.
func (c *Checkbox) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "checkbox",
		Value:   c.label,
		Details: c.prop.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the checkbox will have in the PDF
func (c *Checkbox) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	fontHeight := provider.GetFontHeight(&props.Font{
		Family: c.prop.Family,
		Style:  c.prop.Style,
		Size:   c.prop.Size,
		Color:  c.prop.Color,
	})

	// Height is the maximum of box size and font height, plus padding
	height := c.prop.BoxSize
	if fontHeight > height {
		height = fontHeight
	}

	return height + c.prop.Top + c.prop.Bottom
}

// SetConfig sets the config.
func (c *Checkbox) SetConfig(config *entity.Config) {
	c.config = config
	if c.config != nil {
		c.prop.MakeValid(c.config.DefaultFont)
	}
}

// Render renders a Checkbox into a PDF context.
func (c *Checkbox) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddCheckbox(c.label, cell, &c.prop)
}
