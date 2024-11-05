// rowmapper is the package responsible for mapping row settings
package rowmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type Row struct {
	Height    float64
	Cols      []mappers.Componentmapper
	Factory   mappers.AbstractFactoryMaps
	SourceKey string
}

func NewRow(templateRows interface{}, sourceKey string, factory mappers.AbstractFactoryMaps) (*Row, error) {
	mapRows, ok := templateRows.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure that rows can be converted to map[string] interface{}")
	}
	row := &Row{
		Height:    0,
		Cols:      make([]mappers.Componentmapper, 0),
		Factory:   factory,
		SourceKey: sourceKey,
	}

	err := row.addComponents(mapRows)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// addComponents is responsible for adding the row components according to
// the properties informed in the map
func (r *Row) addComponents(mapRows map[string]interface{}) error {
	fieldMappers := r.getFieldMappers()

	for templateName, template := range mapRows {
		mapper, ok := fieldMappers[templateName]
		if !ok {
			return fmt.Errorf("the field %s present in the row cannot be mapped to any valid component", templateName)
		}
		err := mapper(template)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Row) setHeight(template interface{}) error {
	height, ok := template.(float64)
	if !ok {
		return fmt.Errorf("ensure that height can be converted to float64")
	}
	r.Height = height
	return nil
}

func (r *Row) setCols(template interface{}) error {
	cols, ok := template.([]interface{})
	if !ok {
		return fmt.Errorf("ensure that cols can be converted to []interface{}")
	}
	r.Cols = make([]mappers.Componentmapper, len(cols))

	for i, col := range cols {
		newCol, err := r.Factory.NewCol(col)
		if err != nil {
			return err
		}
		r.Cols[i] = newCol
	}

	return nil
}

// getFieldMappers is responsible for defining which methods are responsible for assembling which components.
// To do this, the component name is linked to a function in a Map.
func (r *Row) getFieldMappers() map[string]func(interface{}) error {
	return map[string]func(interface{}) error{
		"height": r.setHeight,
		"cols":   r.setCols,
	}
}

func (r *Row) getRowContent(content map[string]interface{}) (map[string]interface{}, error) {
	rowContent, ok := content[r.SourceKey]
	if !ok {
		return nil, fmt.Errorf("the row needs the source key \"%s\", but it was not found", r.SourceKey)
	}
	if mapRow, ok := rowContent.(map[string]interface{}); ok {
		return mapRow, nil
	}
	return nil, fmt.Errorf("ensure that the contents of the row \"%s\" can be converted to map[string]interface{}", r.SourceKey)
}

func (r *Row) Generate(content map[string]interface{}, provider processorprovider.ProcessorProvider) ([]processorprovider.ProviderComponent, error) {
	rowContent, err := r.getRowContent(content)
	if err != nil {
		return nil, err
	}

	cols := make([]processorprovider.ProviderComponent, 0, len(r.Cols))
	for _, col := range r.Cols {
		newCol, err := col.Generate(rowContent, provider)
		if err != nil {
			return nil, err
		}
		cols = append(cols, newCol...)
	}

	row, err := provider.CreateRow(r.Height, cols...)
	if err != nil {
		return nil, err
	}
	return []processorprovider.ProviderComponent{row}, nil
}
