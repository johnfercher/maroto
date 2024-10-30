package deserializer_test

import (
	"io"
	"os"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	"github.com/stretchr/testify/assert"
)

func TestDeserializer(t *testing.T) {
	t.Run(
		"when Resources array is not empty, it should return a document with the resources & data for each resource",
		func(t *testing.T) {
			input := `
{
  "Resources": [
    {
      "name": "logo",
      "path": "../../../docs/assets/images/logo.png"
    },
    {
      "name": "font",
      "path": "../../../docs/assets/fonts/arial-unicode-ms.ttf"
    }
  ]
}`

			got, err := deserializer.NewJSONDeserializer().Deserialize(input)
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Contains(t, got, "Resources")

			res, ok := got["Resources"].([]interface{})
			assert.True(t, ok)
			assert.Len(t, res, 2)

			logoFile, err := os.Open("../../../docs/assets/images/logo.png")
			if err != nil {
				t.Fatal(err)
			}
			logoBytes, err := io.ReadAll(logoFile)
			if err != nil {
				t.Fatal(err)
			}

			fontFile, err := os.Open("../../../docs/assets/fonts/arial-unicode-ms.ttf")
			if err != nil {
				t.Fatal(err)
			}
			fontBytes, err := io.ReadAll(fontFile)
			if err != nil {
				t.Fatal(err)
			}

			assetBytes := [][]byte{logoBytes, fontBytes}

			for i, obj := range res {
				resource, ok := obj.(map[string]interface{})
				assert.True(t, ok)
				assert.Contains(t, resource, "data")
				assert.IsType(t, resource["data"], []byte{})
				assert.NotNil(t, resource["data"])
				assert.Equal(t, assetBytes[i], resource["data"])
			}
		},
	)
}
