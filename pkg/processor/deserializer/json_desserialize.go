// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
)

type jsonDeserializer struct {
}

func NewJsonDeserialize() *jsonDeserializer {
	return &jsonDeserializer{}
}

// DesserializeTemplate is responsible for transforming a string into a template structure
func (j *jsonDeserializer) DesserializeTemplate(templateJson string) (mappers.Template, error) {
	return deserializer[mappers.Template](templateJson)
}

// DesserializeContent is responsible for transforming a string into a content structure
func (j *jsonDeserializer) DesserializeContent(contentJson string) (mappers.Content, error) {
	return deserializer[mappers.Content](contentJson)
}

func deserializer[T interface{}](jsonDocument string) (T, error) {
	var document T
	err := json.Unmarshal([]byte(jsonDocument), &document)
	return document, err
}
