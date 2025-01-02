// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"
)

type jsonDeserializer struct{}

// The new JSONserializer is responsible for creating a json deserializer
func NewJSONDeserializer() *jsonDeserializer {
	return &jsonDeserializer{}
}

// Deserialize is responsible for parsing a json document and creating an interface map
func (j *jsonDeserializer) Deserialize(documentJSON string) (map[string]interface{}, error) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(documentJSON), &document)
	return document, err
}
