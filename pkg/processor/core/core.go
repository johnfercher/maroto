package core

type Processor interface {
	RegisterTemplate(templateName string, template string) error
	GenerateDocument(templateName string, content string) []byte
}

type Repository interface {
	RegisterTemplate(name string, template string) error
	ReadTemplate(templateName string) (string, error)
}

type Deserializer interface {
	Deserialize(document string) (map[string]interface{}, error)
}
