package list_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	t.Run("When header is null, should set the header to an empty row", func(t *testing.T) {
		myList := list.New(nil)

		assert.IsType(t, &row.Row{}, myList.Header)
	})
	t.Run("When content is null, should set content to an empty list", func(t *testing.T) {
		myList := list.New(row.New(10))

		assert.Equal(t, 0, len(myList.Content))
	})
}

func TestGetRows(t *testing.T) {
	t.Run("When the list has 3 rows , should return 3 rows", func(t *testing.T) {
		myList := list.New(row.New(10)).Add(row.New(10), row.New(10))

		assert.Equal(t, 3, len(myList.GetRows()))
	})
}

func TestAdd(t *testing.T) {
	t.Run("When null is sent, should not add rows to the list", func(t *testing.T) {
		myList := list.New(row.New(10))

		myList.Add()

		assert.Equal(t, 1, len(myList.GetRows()))
	})
	t.Run("When 10 rows are sent, should add 10 rows to the list", func(t *testing.T) {
		myList := list.New(row.New(10))

		for i := 0; i < 10; i++ {
			myList.Add(row.New(10))
		}

		assert.Equal(t, 10, len(myList.Content))
	})
}
