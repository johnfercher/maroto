package list_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

type anyType struct {
	Key   string
	Value string
}

func (a anyType) GetHeader() core.Row {
	r := row.New(10).Add(
		text.NewCol(6, "Key"),
		text.NewCol(6, "Value"),
	)

	return r
}

func (a anyType) GetContent(i int) core.Row {
	r := row.New(10).Add(
		text.NewCol(6, a.Key),
		text.NewCol(6, a.Value),
	)

	if i%2 == 0 {
		cell := fixture.CellProp()
		r.WithStyle(&cell)
	}

	return r
}

func TestBuild(t *testing.T) {
	t.Run("when arr is empty, should return error", func(t *testing.T) {
		// Act
		r, err := list.Build[anyType](nil)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, r)
	})
	t.Run("when arr is not empty, should return rows", func(t *testing.T) {
		// Arrange
		arr := buildList(10)

		// Act
		r, err := list.Build(arr)
		p := page.New().Add(r...)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(p.GetStructure()).Equals("components/list/build.json")
	})
}

func TestBuildFromPointer(t *testing.T) {
	t.Run("when arr is empty, should return error", func(t *testing.T) {
		// Arrange
		arr := buildPointerList(0)

		// Act
		r, err := list.BuildFromPointer(arr)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, r)
	})
	t.Run("when arr is not empty, should return rows", func(t *testing.T) {
		// Arrange
		arr := buildPointerList(10)

		// Act
		r, _ := list.BuildFromPointer(arr)
		p := page.New().Add(r...)

		// Assert
		test.New(t).Assert(p.GetStructure()).Equals("components/list/build_from_pointer.json")
	})
	t.Run("when arr is has a nil element, should return error", func(t *testing.T) {
		// Arrange
		arr := buildPointerList(10)
		arr[5] = nil

		// Act
		r, err := list.BuildFromPointer(arr)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, r)
	})
}

func buildList(qtd int) []anyType {
	var arr []anyType

	for i := 0; i < qtd; i++ {
		arr = append(arr, anyType{
			Key:   fmt.Sprintf("key(%d)", i),
			Value: fmt.Sprintf("value(%d)", i),
		})
	}

	return arr
}

func buildPointerList(qtd int) []*anyType {
	var arr []*anyType

	for i := 0; i < qtd; i++ {
		arr = append(arr, &anyType{
			Key:   fmt.Sprintf("key(%d)", i),
			Value: fmt.Sprintf("value(%d)", i),
		})
	}

	return arr
}
