package buildermapper

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/test"
	"github.com/stretchr/testify/assert"
)

func DefaultBuilderMap() *Builder {
	time, _ := time.Parse("2006-01-02 15:04:05", "2024-10-09 14:30:00")
	return &Builder{
		Dimensions: &propsmapper.Dimensions{
			Width:  10.0,
			Height: 10.0,
		},
		Margins: &propsmapper.Margins{
			Left:   10.0,
			Right:  10.0,
			Top:    10.0,
			Bottom: 10.0,
		},
		SequentialMode:          false,
		ConcurrentMode:          10,
		SequentialLowMemoryMode: -1,
		Debug:                   true,
		MaxGridSize:             10,
		DefaultFont: &propsmapper.Font{
			Family: "Arial",
			Style:  "bold",
			Size:   10,
			Color: &propsmapper.Color{
				Red:   10,
				Green: 100,
				Blue:  150,
			},
		},
		PageNumber: &propsmapper.PageNumber{
			Pattern: "pattern_test",
			Place:   "place_test",
			Family:  "family_test",
			Style:   "style_test",
			Size:    10.0,
			Color: &propsmapper.Color{
				Red:   10,
				Green: 100,
				Blue:  150,
			},
		},
		CustomFonts: []*propsmapper.CustomFont{
			{Family: "family_test", Style: "style_test", File: "file_test"},
			{Family: "family_test2", Style: "style_test2", File: "file_test2"},
		},
		Protection: &propsmapper.Protection{
			Type:          4,
			UserPassword:  "senha123",
			OwnerPassword: "senha123",
		},
		Compression: true,
		PageSize:    "T",
		Orientation: "vertical",
		Metadata: &propsmapper.Metadata{
			Author:       &propsmapper.Utf8Text{Text: "user_test", UTF8: true},
			Creator:      &propsmapper.Utf8Text{Text: "user_test", UTF8: true},
			Subject:      &propsmapper.Utf8Text{Text: "test", UTF8: true},
			Title:        &propsmapper.Utf8Text{Text: "report", UTF8: true},
			CreationDate: &time,
			KeywordsStr:  &propsmapper.Utf8Text{Text: "test", UTF8: true},
		},
		DisableAutoPageBreak: true,
		GenerationMode:       "concurrent",
	}
}

func TestNewBuilder(t *testing.T) {
	t.Run("when all builder properties are sent, it should generate the builder with all props", func(t *testing.T) {
		builderMap := DefaultBuilderMap()
		file, _ := test.NewFileReader().LoadFile("processor/all_builder_pros.json")

		var builderInterface interface{}
		if err := json.Unmarshal(file, &builderInterface); err != nil {
			t.Error("could not deserialize json")
			return
		}

		generateBuilder, err := NewBuilder(builderInterface)
		assert.Nil(t, err)
		assert.Equal(t, *builderMap, *generateBuilder)
	})

	t.Run("when props is not sent, it should use deafault props", func(t *testing.T) {
		file, _ := test.NewFileReader().LoadFile("processor/without_all_builder_props.json")
		const defaultDimensions = -1.0

		var builderInterface interface{}
		if err := json.Unmarshal(file, &builderInterface); err != nil {
			t.Error("could not deserialize json")
			return
		}

		generateBuilder, err := NewBuilder(builderInterface)
		assert.Nil(t, err)
		assert.Equal(t, defaultDimensions, generateBuilder.Dimensions.Width)
	})

	t.Run("when no builder props is sent, no error is returned", func(t *testing.T) {
		file, _ := test.NewFileReader().LoadFile("processor/without_builder_props.json")

		var builderInterface interface{}
		if err := json.Unmarshal(file, &builderInterface); err != nil {
			t.Error("could not deserialize json")
			return
		}
		_, err := NewBuilder(builderInterface)
		assert.Nil(t, err)
	})

	t.Run("when an invalid Builder is sent, it should return an error", func(t *testing.T) {
		var builderInterface interface{} = 1

		_, err := NewBuilder(builderInterface)
		assert.NotNil(t, err)
	})
}
