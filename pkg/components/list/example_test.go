package list_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// ExampleNew demonstrates how create a list.
func ExampleNew() {
	m := maroto.New()

	myList := list.New(text.NewAutoRow("Header"))
	myList.Add(text.NewAutoRow("Content_1"), text.NewAutoRow("Content_2"))

	m.AddRows(myList.GetRows()...)

	// generate document
}

// ExampleNew demonstrates how to add content to the list.
func ExampleList_Add() {
	m := maroto.New()

	myList := list.New(text.NewAutoRow("Header")).Add(text.NewAutoRow("Content"), text.NewAutoRow("Content"))

	m.AddRows(myList.GetRows()...)

	// generate document
}

func ExampleList_BuildListWithFixedHeader() {
	m := maroto.New()

	myList := list.New(text.NewAutoRow("Header"), props.List{MinimumRowsBypage: 2})
	myList.Add(text.NewAutoRow("Content_1"), text.NewAutoRow("Content_2"))

	_ = myList.BuildListWithFixedHeader(m)

	// generate document
}
