package core

type Processor interface {
	RegisterTemplate(templateName string, template string) error
	GenerateDocument(templateName string, content string) []byte
}

type Repository interface {
	RegisterTemplate(name string, template string) error
	ReadTemplate(templateName string) (string, error)
}

type DocumentDeserializer[T interface{}] interface {
	DesserializeTemplate(template string) (T, error)
	DesserializeContent(content string) (map[string]interface{}, error)
}

type Component interface {
}

type Provider interface {
	GeneratePdf(componentTree Component) ([]byte, error)
}
