package processor

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/abstractfactory"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/documentmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/repository"
)

type processor struct {
	repository   core.Repository
	deserializer core.Deserializer
	loader       core.Loader
}

// NewProcessor is responsible for creating a processor object
// The processor object should be used to create PDF from a serialized document
func NewProcessor() *processor {
	return &processor{
		repository:   repository.NewMemoryStorage(),
		deserializer: deserializer.NewJSONDeserializer(),
	}
}

// RegisterTemplate is responsible for recording the document template.
// Each template must be accompanied by a name that will identify it.
func (p *processor) RegisterTemplate(templateName string, template string) error {
	t, err := p.deserializer.Deserialize(template)
	if err != nil {
		return err
	}
	return p.repository.RegisterTemplate(templateName, t)
}

// GenerateDocument is responsible for generating the pdf
// templateName must reference a previously saved template,
// this template will be computed with the data sent in content
func (p *processor) GenerateDocument(templateName string, content string) ([]byte, error) {
	template, err := p.repository.ReadTemplate(templateName)
	if err != nil {
		return nil, err
	}

	document, err := documentmapper.NewPdf(template, abstractfactory.NewAbstractFactoryMaps(p.repository))
	if err != nil {
		return nil, err
	}
	fmt.Print(document)

	return nil, nil
}
