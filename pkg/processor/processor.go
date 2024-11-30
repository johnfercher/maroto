package processor

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/abstractfactory"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/documentmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type processor struct {
	repository   core.ProcessorRepository
	deserializer core.Deserializer
}

// NewProcessor is responsible for creating a processor object
// The processor object should be used to create PDF from a serialized document
func NewProcessor(repository core.ProcessorRepository, deserializer core.Deserializer) *processor {
	return &processor{
		repository:   repository,
		deserializer: deserializer,
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
func (p *processor) GenerateDocument(templateName string, content string) (*processorprovider.ProcessorProvider, error) {
	template, err := p.repository.ReadTemplate(templateName)
	if err != nil {
		return nil, err
	}

	document, err := documentmapper.NewPdf(template, abstractfactory.NewAbstractFactoryMaps(p.repository))
	if err != nil {
		return nil, err
	}

	marotoProvider, err := processorprovider.NewMaroto(p.repository, *document.GetBuilderCfg())
	if err != nil {
		return nil, err
	}

	contentMap, err := p.deserializer.Deserialize(content)
	if err != nil {
		return nil, err
	}

	provider, err := (*document).Generate(contentMap, marotoProvider)
	if err != nil {
		return nil, err
	}
	return provider, nil
}
