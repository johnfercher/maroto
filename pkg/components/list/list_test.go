package list_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
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

func TestBuildListWithFixedHeader(t *testing.T) {
	t.Run("when Maroto is not sent, an error is returned", func(t *testing.T) {
		myList := list.New(row.New(10)).Add(row.New(10))

		err := myList.BuildListWithFixedHeader(nil)

		assert.NotNil(t, err)
	})
	t.Run("when list uses 2 pages, should repeat the header on both pages", func(t *testing.T) {
		m := mocks.NewMaroto(t)

		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(2)
		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content2")).Return(2)

		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content1"))
		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content2"))

		m.EXPECT().FillPageToAddNew()

		myList := list.New(text.NewAutoRow("header")).Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"))

		assert.Nil(t, myList.BuildListWithFixedHeader(m))

		m.AssertNumberOfCalls(t, "AddRows", 2)
		m.AssertNumberOfCalls(t, "FitsOnCurrentPage", 2)
		m.AssertNumberOfCalls(t, "FillPageToAddNew", 1)
	})

	t.Run("when the list has auto-height rows, it should generate the list", func(t *testing.T) {
		// Arrange
		sut := maroto.New()
		// Act
		myList := list.New(text.NewAutoRow("header")).Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"))
		err := myList.BuildListWithFixedHeader(sut)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("components/list/list_with_auto_row.json")
	})

	// nolint: dupl
	t.Run("when the header does not fit on the current page, should move the list to the next page", func(t *testing.T) {
		m := mocks.NewMaroto(t)

		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(0).Once()
		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(3)

		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2"))
		m.EXPECT().FillPageToAddNew()

		myList := list.New(text.NewAutoRow("header"), props.List{MinimumRowsBypage: 2})
		myList.Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"))

		assert.Nil(t, myList.BuildListWithFixedHeader(m))

		m.AssertNumberOfCalls(t, "AddRows", 1)
		m.AssertNumberOfCalls(t, "FitsOnCurrentPage", 2)
		m.AssertNumberOfCalls(t, "FillPageToAddNew", 1)
	})
	// nolint: dupl
	t.Run("when only header fits on current page, should move list to next page", func(t *testing.T) {
		m := mocks.NewMaroto(t)

		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(1).Once()
		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(3)

		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2"))
		m.EXPECT().FillPageToAddNew()

		myList := list.New(text.NewAutoRow("header"), props.List{MinimumRowsBypage: 2})
		myList.Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"))

		assert.Nil(t, myList.BuildListWithFixedHeader(m))

		m.AssertNumberOfCalls(t, "AddRows", 1)
		m.AssertNumberOfCalls(t, "FitsOnCurrentPage", 2)
		m.AssertNumberOfCalls(t, "FillPageToAddNew", 1)
	})

	t.Run("when list fit on current page, should add list to the current page", func(t *testing.T) {
		m := mocks.NewMaroto(t)

		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2")).Return(3)
		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2"))

		myList := list.New(text.NewAutoRow("header")).Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"))

		assert.Nil(t, myList.BuildListWithFixedHeader(m))

		m.AssertNumberOfCalls(t, "AddRows", 1)
		m.AssertNumberOfCalls(t, "FitsOnCurrentPage", 1)
	})

	t.Run("when it is not possible to add a group of rows, should return error", func(t *testing.T) {
		m := mocks.NewMaroto(t)

		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content1"),
			text.NewAutoRow("content2"), text.NewAutoRow("content3")).Return(3)
		m.EXPECT().FitsOnCurrentPage(text.NewAutoRow("header"), text.NewAutoRow("content3")).Return(0)
		m.EXPECT().AddRows(text.NewAutoRow("header"), text.NewAutoRow("content1"), text.NewAutoRow("content2"))
		m.EXPECT().FillPageToAddNew()

		myList := list.New(text.NewAutoRow("header"))
		myList.Add(text.NewAutoRow("content1"), text.NewAutoRow("content2"), text.NewAutoRow("content3"))

		assert.NotNil(t, myList.BuildListWithFixedHeader(m))
		m.AssertNumberOfCalls(t, "AddRows", 1)
		m.AssertNumberOfCalls(t, "FitsOnCurrentPage", 3)
		m.AssertNumberOfCalls(t, "FillPageToAddNew", 2)
	})
}
