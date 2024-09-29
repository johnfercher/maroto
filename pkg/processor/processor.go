package processor

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/repository"
)

type processor struct {
	repository   core.Repository
	deserializer core.DocumentDeserializer
	provider     core.Provider
}

func NewProcessor() *processor {
	return &processor{repository: repository.NewMemoryStorage()}
}

func (p *processor) RegisterTemplate(templateName string, template string) error {
	return p.repository.RegisterTemplate(templateName, template)
}

func (p *processor) GenerateDocument(templateName string, content string) ([]byte, error) {
	templateJson, err := p.repository.ReadTemplate(templateName)
	if err != nil {
		return nil, err
	}

	documentTemplate, err := p.deserializer.DesserializeTemplate(templateJson)
	if err != nil {
		return nil, err
	}

	documentContent, err := p.deserializer.DesserializeContent(templateJson)
	if err != nil {
		return nil, err
	}

	pdfComponent, err := documentTemplate.Generate(documentContent)
	if err != nil {
		return nil, err
	}

	return p.provider.GeneratePdf(pdfComponent)
}
