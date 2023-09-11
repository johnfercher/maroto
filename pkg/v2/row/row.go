package row

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v2"
)

type row struct {
	height     int
	_type      string
	accept     map[string]bool
	components []v2.Component
}

func New(height int) *row {
	accept := make(map[string]bool)
	accept[v2.Col] = true

	return &row{
		_type:  v2.Row,
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

func (d *row) Add(components ...v2.Component) {
	for _, component := range components {
		if _, ok := d.accept[component.GetType()]; ok {
			d.components = append(d.components, component)
		}
	}
}
