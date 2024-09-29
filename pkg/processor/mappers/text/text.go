// text is the package responsible for mapping text settings
package text

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/props"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/text"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Text struct {
	Props        propsmapper.TextProps `json:"props"`
	SourceKey    string                `json:"source_key"`
	DefaultValue string                `json:"value"`
}

// generate is responsible for the builder text according to the submitted content
func (t *Text) Generate(content map[string]interface{}) (components.Component, error) {
	if t.DefaultValue != "" {
		return text.NewText(props.TextProps{Align: t.Props.Align}, t.DefaultValue), nil
	}

	value, ok := content[t.SourceKey]
	if !ok {
		return nil, fmt.Errorf("text model needs source key %s, but no content with that key was found", t.SourceKey)
	}

	textValue, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("resource %s does not have a valid value for the text component", t.SourceKey)
	}

	return text.NewText(props.TextProps{Align: t.Props.Align}, textValue), nil
}
