package main

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	t.Parallel()
	// Act
	sut := GetMaroto()

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/custompage.json")
}
