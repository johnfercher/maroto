package colmapper_test

import (
	"errors"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/colmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/imagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/linemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/signaturemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/textmapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCol(t *testing.T) {
	t.Run("when a barcode is sent, a barcode is created", func(t *testing.T) {
		col := map[string]interface{}{"bar_code": nil}
		validBarcode := fixture.Barcode()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewBarcode", mock.Anything).Return(validBarcode, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &codemapper.Barcode{}, doc.Components[0])
	})
	t.Run("when a matrixcode is sent, a matrixcode is created", func(t *testing.T) {
		col := map[string]interface{}{"matrix_code": nil}
		validMatrixcode := fixture.Matrixcode()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewMatrixcode", mock.Anything).Return(validMatrixcode, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &codemapper.Matrixcode{}, doc.Components[0])
	})
	t.Run("when a qrcode is sent, a qrcode is created", func(t *testing.T) {
		col := map[string]interface{}{"qr_code": nil}
		validQrcode := fixture.Qrcode()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewQrcode", mock.Anything).Return(validQrcode, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &codemapper.Qrcode{}, doc.Components[0])
	})
	t.Run("when a image is sent, a image is created", func(t *testing.T) {
		col := map[string]interface{}{"image": nil}
		validImage := fixture.Image()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewImage", mock.Anything).Return(validImage, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &imagemapper.Image{}, doc.Components[0])
	})
	t.Run("when a line is sent, a line is created", func(t *testing.T) {
		col := map[string]interface{}{"line": nil}
		validLine := fixture.Line()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewLine", mock.Anything).Return(validLine, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &linemapper.Line{}, doc.Components[0])
	})
	t.Run("when a signature is sent, a signature is created", func(t *testing.T) {
		col := map[string]interface{}{"signature": nil}
		validSignature := fixture.Signature()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewSignature", mock.Anything).Return(validSignature, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &signaturemapper.Signature{}, doc.Components[0])
	})
	t.Run("when a text is sent, a text is created", func(t *testing.T) {
		col := map[string]interface{}{"text": nil}
		validText := fixture.Text()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewText", mock.Anything).Return(validText, nil)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 1)
		assert.IsType(t, &textmapper.Text{}, doc.Components[0])
	})
	t.Run("when no component is sent, no component is added", func(t *testing.T) {
		col := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Len(t, doc.Components, 0)
	})
	t.Run("when an invalid field is sent, an error is returned", func(t *testing.T) {
		col := map[string]interface{}{"invalid_field": nil}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when an invalid interface is sent, an error is returned", func(t *testing.T) {
		var col interface{} = 1
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when the component cannot be created, it should return an error", func(t *testing.T) {
		col := map[string]interface{}{"text": nil}
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewText", mock.Anything).Return(nil, errors.New("any"))

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when an invalid size is sent, an error is returned", func(t *testing.T) {
		col := map[string]interface{}{"size": "invalid"}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when no size is sent, should set size to 0", func(t *testing.T) {
		col := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Equal(t, doc.Size, 0)
	})
	t.Run("when size is sent, should set size", func(t *testing.T) {
		col := map[string]interface{}{"size": 6.0}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := colmapper.NewCol(col, factory)

		assert.Nil(t, err)
		assert.Equal(t, doc.Size, 6)
	})
}
