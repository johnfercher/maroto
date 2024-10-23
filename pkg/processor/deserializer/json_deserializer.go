// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/loader"
)

type jsonDeserializer struct {
	loader core.Loader
}

func NewJsonDeserializer() *jsonDeserializer {
	return &jsonDeserializer{loader: loader.NewLoader()}
}

func (j *jsonDeserializer) Deserialize(documentJson string) (map[string]interface{}, error) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(documentJson), &document)
	return document, err
}
