package core

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pdfmapper"
)

type Processor interface {
	RegisterTemplate(templateName string, template string) error
	GenerateDocument(templateName string, content string) []byte
}

type Repository interface {
	RegisterTemplate(name string, template string) error
	ReadTemplate(templateName string) (string, error)
}

type DocumentDeserializer interface {
	DesserializeTemplate(template string) (pdfmapper.Pdf, error)
	DesserializeContent(content string) (map[string]interface{}, error)
}
