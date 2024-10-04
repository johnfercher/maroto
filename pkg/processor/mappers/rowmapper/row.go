// rowmapper is the package responsible for mapping row settings
package rowmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components/col"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/row"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/colmapper"
)

type Row struct {
	List string          `json:"list"`
	Cols []colmapper.Col `json:"cols"`
}

// generate is responsible for the builder row according to the submitted content
func (r *Row) Generate(content map[string]interface{}) (*row.Row, error) {
	if len(r.List) > 0 {
		listContent, err := r.getListContent(r.List, content)
		if err != nil {
			return nil, err
		}
		content = listContent
	}

	cols, err := r.generateCols(content)
	if err != nil {
		return nil, err
	}
	return row.NewRow(cols...), nil
}

func (r *Row) getListContent(listKey string, content map[string]interface{}) (map[string]interface{}, error) {
	contentList, ok := content[listKey]
	if !ok {
		return nil, fmt.Errorf("the model needed a list with key %s, but that key was not found in the content", r.List)
	}

	contentMap, ok := contentList.([]interface{})
	if !ok {
		return nil, fmt.Errorf("key \"%s\" references a content that cannot be converted to a valid format, ensure that this content can be converted to a map[string]interface{}", r.List)
	}

	b := contentMap[0]
	c, ok := b.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("key \"%s\" references a content that cannot be converted to a valid format, ensure that this content can be converted to a map[string]interface{}", r.List)
	}

	return c, nil
}

func (r *Row) generateCols(content map[string]interface{}) ([]col.Col, error) {
	generatedCols := make([]col.Col, len(r.Cols))

	for i, col := range r.Cols {
		generatedCol, err := col.Generate(content)
		if err != nil {
			return nil, err
		}
		generatedCols[i] = *generatedCol
	}
	return generatedCols, nil
}
