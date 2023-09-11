package col

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v22"
)

type col struct {
	size       int
	_type      string
	accept     map[string]bool
	components []v22.Component
}

func New(size int) *col {
	accept := make(map[string]bool)
	accept[v22.Image] = true

	return &col{
		_type:  v22.Col,
		accept: accept,
		size:   size,
	}
}

func (d *col) Render() {
	fmt.Println(d.size)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *col) GetType() string {
	return d._type
}

func (d *col) Add(components ...v22.Component) {
	for _, component := range components {
		if _, ok := d.accept[component.GetType()]; ok {
			d.components = append(d.components, component)
		}
	}
}
