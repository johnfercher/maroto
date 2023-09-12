package col

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v2"
)

type col struct {
	size       int
	_type      v2.DocumentType
	components []v2.Component
}

func New(size int) *col {
	return &col{
		_type: v2.Col,
		size:  size,
	}
}

func (d *col) Render() {
	fmt.Println(d.size)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *col) GetType() string {
	return d._type.String()
}

func (d *col) Add(components ...v2.Component) {
	for _, component := range components {
		if d._type.Accept(component.GetType()) {
			d.components = append(d.components, component)
		}
	}
}
