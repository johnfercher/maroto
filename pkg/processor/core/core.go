package core

type Processor interface {
	RegisterTemplate(templateName string, template string) error
	GenerateDocument(templateName string, content string) []byte
}

type Repository interface {
	RegisterTemplate(templateName string, template map[string]any) error
	ReadTemplate(templateName string) (map[string]any, error)
}

type Deserializer interface {
	Deserialize(document string) (map[string]interface{}, error)
}

