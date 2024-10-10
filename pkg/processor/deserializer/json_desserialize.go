// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"
)

type jsonDeserializer struct{}

func NewJsonDeserialize() *jsonDeserializer {
	return &jsonDeserializer{}
}

func (j *jsonDeserializer) Deserialize(documentJson string) (map[string]interface{}, error) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(documentJson), &document)
	return document, err
}
