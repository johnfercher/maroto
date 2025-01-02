package main

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	test.SetupTestDir(t)
	// Act
	sut := GetMaroto()

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/imagegrid.json")
}
