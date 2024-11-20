// The deserialize package is responsible for assembling the structures used in the processor according to the receiving string.
package deserializer

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/loader"
)

type jsonDeserializer struct {
	loader core.Loader
}

// The new JSONserializer is responsible for creating a json deserializer
func NewJSONDeserializer() *jsonDeserializer {
	return &jsonDeserializer{loader: loader.NewLoader()}
}

// Deserialize is responsible for parsing a json document and creating an interface map
func (j *jsonDeserializer) Deserialize(documentJSON string) (map[string]interface{}, error) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(documentJSON), &document)
	if err != nil {
		return nil, err
	}

	resources, ok := document["Resources"].([]interface{})
	if !ok {
		return document, nil
	}

	document["Resources"] = j.loadResources(resources)

	return document, nil
}

// loadResources method is responsible for loading each resource into
// memory via go routines
// `resources` should be type []map[string]interface{} with key "path"
// returns []map[string]interface{} with key "data" and the associated bytes of the file
func (j *jsonDeserializer) loadResources(resources []interface{}) []interface{} {
	wg := sync.WaitGroup{}
	for i, res := range resources {
		resource, ok := res.(map[string]interface{})
		if !ok {
			continue
		}
		path, ok := resource["path"].(string)
		if !ok {
			continue
		}

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data, err := j.loader.Load(path)
			if err != nil {
				// TODO handle errors && io gracefully
				fmt.Fprintln(os.Stderr, err.Error())
			}
			resource["data"] = data
			resources[i] = resource
		}(i)
	}
	wg.Wait()
	return resources
}
