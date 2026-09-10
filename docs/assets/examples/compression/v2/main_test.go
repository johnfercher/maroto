package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	t.Parallel()
	// Act
	path := "docs/assets/images/frontpage.png"
	sut := GetMaroto(buildPath(path))

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/compression.json")
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = filepath.Join(dir, "..", "..", "..", "..", "..")
	return filepath.Join(dir, file)
}
