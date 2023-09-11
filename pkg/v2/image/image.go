package image

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/v2"
)

type image struct {
	path       string
	_type      string
	components []v2.Component
}

func New(path string) *image {
	return &image{
		_type: v2.Image,
		path:  path,
	}
}

func (d *image) Render() {
	fmt.Println(d.path)
	for _, component := range d.components {
		component.Render()
	}
}

func (d *image) GetType() string {
	return d._type
}

func (d *image) Add(_ ...v2.Component) {
	return
}
