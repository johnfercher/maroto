package core

type Processor interface {
	RegisterTemplate(template string) error
	GenerateDocument(templateId int, content string) []byte
}

type Repository interface {
	RegisterTemplate(template string) error
	ReadTemplate(templateId int) (string, error)
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
