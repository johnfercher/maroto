package core

import (
	"html/template"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/content"
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
	DesserializeTemplate(template string) (template.Template, error)
	DesserializeContent(content string) (content.Content, error)
}

type Component interface {
}

type Provider interface {
	GeneratePdf(componentTree Component) ([]byte, error)
}
