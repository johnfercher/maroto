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
}

func NewProcessor() *processor {
	return &processor{repository: repository.NewMemoryStorage(), deserializer: deserializer.NewJsonDeserialize()}
}

func (p *processor) RegisterTemplate(templateName string, template string) error {
	return p.repository.RegisterTemplate(templateName, template)
}

func (p *processor) GenerateDocument(templateName string, content string) ([]byte, error) {
	templateJson, err := p.repository.ReadTemplate(templateName)
	if err != nil {
		return nil, err
	}

	document, err := documentmapper.NewPdf(templateJson, p.deserializer, abstractfactory.NewAbstractFactoryMaps())
	if err != nil {
		return nil, err
	}
	fmt.Print(document)

	return nil, nil
}
