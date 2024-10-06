package processor

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
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

	documentTemplate, err := p.deserializer.Deserialize(templateJson)
	if err != nil {
		return nil, err
	}
	fmt.Println(documentTemplate)

	// documentContent, err := p.deserializer.DesserializeContent(content)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
