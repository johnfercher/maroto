package core

type Processor interface {
	RegisterTemplate(templateName string, template string) error
	GenerateDocument(templateName string, content string) []byte
}

type ProcessorRepository interface {
	RegisterTemplate(templateName string, template map[string]any) error
	ReadTemplate(templateName string) (map[string]any, error)
	GetDocument(documentName string) (extension string, doc []byte, err error)
}

type Deserializer interface {
	Deserialize(document string) (map[string]interface{}, error)
}

// Takes a path and returns its bytes
// path may be file path or url
type Loader interface {
	Load(path string) ([]byte, error)
}
