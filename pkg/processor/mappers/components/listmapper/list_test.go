package listmapper_test

import (
	"errors"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewList(t *testing.T) {
	t.Run("when invalid interface is sent, it should return an error", func(t *testing.T) {
		var invalidInterface interface{} = 1
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := listmapper.NewList(invalidInterface, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when component not can generate, it should return an error", func(t *testing.T) {
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
		}

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(nil, errors.New(""))
		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when 2-components are sent, it should add 2 components in list", func(t *testing.T) {
		validPage := fixture.MapperPage()
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
		}
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(validPage, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(validPage, nil)

		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Templates), 2)
	})
}
