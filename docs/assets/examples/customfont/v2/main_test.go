package main

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/test"
)

func TestGetMaroto(t *testing.T) {
	// Act
	sut := GetMaroto(buildPath("docs/assets/fonts/arial-unicode-ms.ttf"))

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/customfont.json")
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "docs/assets/examples/customfont/v2", "")
	return path.Join(dir, file)
}
