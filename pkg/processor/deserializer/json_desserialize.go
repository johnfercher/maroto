// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"
	"html/template"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/content"
)

type jsonDeserializer struct {
}

func NewJsonDeserialize() *jsonDeserializer {
	return &jsonDeserializer{}
}

// DesserializeTemplate is responsible for transforming a string into a template structure
func (j *jsonDeserializer) DesserializeTemplate(templateJson string) (template.Template, error) {
	return deserializer[template.Template](templateJson)
}

// DesserializeContent is responsible for transforming a string into a content structure
func (j *jsonDeserializer) DesserializeContent(contentJson string) (content.Content, error) {
	return deserializer[content.Content](contentJson)
}

func deserializer[T interface{}](jsonDocument string) (T, error) {
	var document T
	err := json.Unmarshal([]byte(jsonDocument), &document)
	return document, err
}
