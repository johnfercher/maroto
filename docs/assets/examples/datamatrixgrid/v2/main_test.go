package main

import (
	"github.com/johnfercher/maroto/v2/pkg/test"
	"testing"
)

func TestGetMaroto(t *testing.T) {
	// Act
	sut := GetMaroto()

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/datamatrixgrid.json")
}
