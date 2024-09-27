package mappers

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/foreach"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pages"
)

// Template defines the structure that should be at the top level of the template.
// It allows you to add PDF settings and add a list of pages.
type Template struct {
	// The builder defines the PDF settings
	Builder builder.Builder

	// ForEach defines the structure that will allow you to add a list of pages to the template
	ForEach foreach.ForEach[pages.Page]
}
