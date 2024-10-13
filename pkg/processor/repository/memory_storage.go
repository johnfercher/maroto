// repository package is responsible for managing access to templates
package repository

type memoryStorage struct {
	template map[string]string
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		template: make(map[string]string),
	}
}

// RegisterTemplate is responsible for register a template in memory
//   - name is the model identifier and is used to access it
//   - template is the template that will be stored
func (m *memoryStorage) RegisterTemplate(name string, template string) error {
	m.template[name] = template
	return nil
}

// ReadTemplate is responsible for fetching the stored template
//   - name is the model identifier and is used to access it
func (m *memoryStorage) ReadTemplate(templateName string) (string, error) {
	return m.template[templateName], nil
}