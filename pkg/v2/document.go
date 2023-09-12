package v2

import "fmt"

type document struct {
	value      string
	_type      DocumentType
	components []Component
}

func NewDocument(value string) *document {
	return &document{
		_type: Document,
		value: value,
	}
}

func (d *document) Render() {
	fmt.Println(d.value)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *document) IsDrawable() bool {
	return false
}

func (d *document) GetType() string {
	return d._type.String()
}

func (d *document) Add(components ...Component) {
	for _, component := range components {
		if d._type.Accept(component.GetType()) {
			d.components = append(d.components, component)
		}
	}
}
