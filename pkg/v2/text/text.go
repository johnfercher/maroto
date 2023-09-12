package text

import (
	"fmt"
	v2 "github.com/johnfercher/maroto/pkg/v2"
)

type text struct {
	value      string
	_type      v2.DocumentType
	components []v2.Component
}

func New(value string) *text {
	return &text{
		_type: v2.Text,
		value: value,
	}
}

func (d *text) Render() {
	fmt.Println(d.value)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *text) GetType() string {
	return d._type.String()
}

func (d *text) Add(_ ...v2.Component) {
	return
}
