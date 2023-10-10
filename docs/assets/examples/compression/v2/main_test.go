package main

import (
	"github.com/johnfercher/maroto/v2/pkg/test"
	"os"
	"path"
	"strings"
	"testing"
)

func TestGetMaroto(t *testing.T) {
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

	dir = strings.ReplaceAll(dir, "docs/assets/examples/compression/v2", "")
	return path.Join(dir, file)
}
