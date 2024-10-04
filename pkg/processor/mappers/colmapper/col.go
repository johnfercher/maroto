// colmapper is the package responsible for mapping col settings
package colmapper

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/col"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/barcode"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/image"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/text"
)

type Col struct {
	Props   propsmapper.ColProps `json:"props"`
	Text    []text.Text          `json:"text"`
	BarCode []barcode.BarCode    `json:"bar_code"`
	Image   []image.Image        `json:"image"`
}

// generate is responsible for the builder col according to the submitted content
func (c *Col) Generate(content map[string]interface{}) (*col.Col, error) {
	components, err := c.factoryComponents(content)
	if err != nil {
		return nil, err
	}
	return col.NewCol(props.ColProps{Size: c.Props.Size}, components...), nil
}

func (c *Col) factoryComponents(content map[string]interface{}) ([]components.Component, error) {
	componentsTemplate := make([]mappers.Componentmapper, 0)

	for _, t := range c.Text {
		componentsTemplate = append(componentsTemplate, &text.Text{Props: t.Props, SourceKey: t.SourceKey, DefaultValue: t.DefaultValue})
	}

	for _, barcode := range c.BarCode {
		componentsTemplate = append(componentsTemplate, &barcode)
	}

	components, err := c.generateComponents(content, componentsTemplate...)
	if err != nil {
		return nil, err
	}

	return components, nil
}

func (c *Col) generateComponents(content map[string]interface{}, templates ...mappers.Componentmapper) ([]components.Component, error) {
	components := make([]components.Component, len(templates))

	if len(templates) == 0 {
		return components, nil
	}
	for i, template := range templates {
		component, err := template.Generate(content)
		if err != nil {
			return nil, err
		}
		components[i] = component
	}
	return components, nil
}
