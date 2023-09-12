package row

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v2"
)

type row struct {
	height     int
	_type      v2.DocumentType
	components []v2.Component
}

func New(height int) *row {
	return &row{
		_type:  v2.Row,
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
	return d._type.String()
}

func (d *row) Add(components ...v2.Component) {
	for _, component := range components {
		if d._type.Accept(component.GetType()) {
			d.components = append(d.components, component)
		}
	}
}
