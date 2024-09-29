package text

import "github.com/johnfercher/maroto/v2/pkg/processor/mappers/props/text"

type Text struct {
	Props     text.TextProps `json:"props"`
	SourceKey string         `json:"source_key"`
	Value     string         `json:"value"`
}
