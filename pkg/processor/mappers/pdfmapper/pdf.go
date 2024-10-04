// padmapper is the package responsible for mapping pdf settings
package pdfmapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/components/builder"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/page"
	"github.com/johnfercher/maroto/v2/pkg/processor/components/pdf"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pagemapper"
)

type Pdf struct {
	Builder buildermapper.Builder      `json:"builder"`
	Pages   map[string]pagemapper.Page `json:"pages"`
}

// generate is responsible for the builder pdf according to the submitted content
func (p *Pdf) Generate(content map[string]interface{}) (*pdf.Pdf, error) {
	var pages []*page.Page

	for pageKey, pageContent := range content {
		pageTemplate, ok := p.Pages[pageKey]
		if !ok {
			return nil, fmt.Errorf("the document content references a page template with key \"%s\", but no page with that key was found in the current template", pageKey)
		}

		content, ok := pageContent.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("key \"%s\" references a content that cannot be converted to a valid format, ensure that this content can be converted to a map[string]interface{}", pageKey)
		}

		generatedPage, err := pageTemplate.Generate(content)
		if err != nil {
			return nil, err
		}

		pages = append(pages, generatedPage)
	}
	return pdf.NewPdf(builder.NewBuilder(), pages...), nil
}
