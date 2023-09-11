package v22

import "fmt"

type document struct {
	value      string
	_type      string
	accept     map[string]bool
	components []Component
}

func NewDocument(value string) *document {
	accept := make(map[string]bool)
	accept[Row] = true

	return &document{
		_type:  Document,
		accept: accept,
		value:  value,
	}
}

func (d *document) Render() {
	fmt.Println(d.value)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *document) GetType() string {
	return d._type
}

func (d *document) Add(components ...Component) {
	for _, component := range components {
		if _, ok := d.accept[component.GetType()]; ok {
			d.components = append(d.components, component)
		}
	}
}
