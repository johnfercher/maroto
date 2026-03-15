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
	prop := props.Checkbox{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &Checkbox{
		label: label,
		prop:  prop,
	}
}

// NewCol is responsible to create an instance of a Checkbox wrapped in a Col.
func NewCol(size int, label string, ps ...props.Checkbox) core.Col {
	c := New(label, ps...)
	return col.New(size).Add(c)
}

// NewAutoRow is responsible for creating an instance of Checkbox grouped in a Row with automatic height.
func NewAutoRow(label string, ps ...props.Checkbox) core.Row {
	c := New(label, ps...)
	cl := col.New().Add(c)
	return row.New().Add(cl)
}

// NewRow is responsible to create an instance of a Checkbox wrapped in a Row.
func NewRow(height float64, label string, ps ...props.Checkbox) core.Row {
	c := New(label, ps...)
	cl := col.New().Add(c)
	return row.New(height).Add(cl)
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

// SetConfig sets the config.
func (c *Checkbox) SetConfig(config *entity.Config) {
	c.config = config
}

// GetHeight returns the height that the checkbox will have in the PDF.
func (c *Checkbox) GetHeight(_ core.Provider, _ *entity.Cell) float64 {
	return c.prop.Size + c.prop.Top
}

// Render renders a Checkbox into a PDF context.
func (c *Checkbox) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddCheckbox(c.label, cell, &c.prop)
}
