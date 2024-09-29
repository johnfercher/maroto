// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pdf"
)

type jsonDeserializer struct {
}

func NewJsonDeserialize() *jsonDeserializer {
	return &jsonDeserializer{}
}

// DesserializeTemplate is responsible for transforming a string into a template structure
func (j *jsonDeserializer) DesserializeTemplate(templateJson string) (pdf.Pdf, error) {
	return deserializer[pdf.Pdf](templateJson)
}

// DesserializeContent is responsible for transforming a string into a content map
func (j *jsonDeserializer) DesserializeContent(contentJson string) (map[string]interface{}, error) {
	return deserializer[map[string]interface{}](contentJson)
}

func deserializer[T interface{}](jsonDocument string) (T, error) {
	var document T
	err := json.Unmarshal([]byte(jsonDocument), &document)
	return document, err
}
