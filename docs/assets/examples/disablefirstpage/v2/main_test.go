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
	path := "docs/assets/images/certificate.png"
	sut := GetMaroto(buildPath(path))

	// Assert
	test.New(t).Assert(sut.GetStructure()).Equals("examples/disablefirstpage.json")
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "docs/assets/examples/disablefirstpage/v2", "")
	return path.Join(dir, file)
}
