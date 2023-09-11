package row

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v22"
)

type row struct {
	height     int
	_type      string
	accept     map[string]bool
	components []v22.Component
}

func New(height int) *row {
	accept := make(map[string]bool)
	accept[v22.Col] = true

	return &row{
		_type:  v22.Row,
		accept: accept,
		height: height,
	}
}

func (d *row) Render() {
	fmt.Println(d.height)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *row) GetType() string {
	return d._type
}

func (d *row) Add(components ...v22.Component) {
	for _, component := range components {
		if _, ok := d.accept[component.GetType()]; ok {
			d.components = append(d.components, component)
		}
	}
}
