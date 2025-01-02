// repository package is responsible for managing access to templates
package repository

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
)

type memoryStorage struct {
	template  map[string]map[string]any
	documents map[string][]byte
	loader    core.Loader
}

// NewMemoryStorage is responsible for creating a repository
// implementation that stores data in memory
func NewMemoryStorage(loader core.Loader) *memoryStorage {
	return &memoryStorage{
		template:  make(map[string]map[string]any),
		loader:    loader,
		documents: make(map[string][]byte),
	}
}

// GetDocument is responsible search and return the document according to the name sent
//   - documentName is the name that the document references
func (m *memoryStorage) GetDocument(documentPath string) (string, []byte, error) {
	if doc, ok := m.documents[documentPath]; ok {
		return m.loader.GetExt(documentPath), doc, nil
	}

	bytes, err := m.loader.Load(documentPath)
	if err != nil {
		return "", nil, err
	}
	m.documents[documentPath] = bytes
	return m.loader.GetExt(documentPath), bytes, nil
}

// RegisterTemplate is responsible for register a template in memory
//   - name is the model identifier and is used to access it
//   - template is the template that will be stored
func (m *memoryStorage) RegisterTemplate(name string, template map[string]any) error {
	m.template[name] = template
	return nil
}

// ReadTemplate is responsible for fetching the stored template
//   - name is the model identifier and is used to access it
func (m *memoryStorage) ReadTemplate(templateName string) (map[string]any, error) {
	return m.template[templateName], nil
}
