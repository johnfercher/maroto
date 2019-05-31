package maroto_test

import (
	"github.com/johnfercher/maroto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFont_GetStyleString(t *testing.T) {
	cases := []struct {
		name        string
		style       maroto.Style
		styleString string
	}{
		{
			"maroto.Normal",
			maroto.Normal,
			"",
		},
		{
			"maroto.Bold",
			maroto.Bold,
			"B",
		},
		{
			"maroto.Italic",
			maroto.Italic,
			"I",
		},
		{
			"maroto.BoldItalic",
			maroto.BoldItalic,
			"BI",
		},
	}

	for _, c := range cases {
		// Act
		styleString := maroto.GetStyleString(c.style)

		// Assert
		assert.Equal(t, styleString, c.styleString)
	}
}

func TestFont_GetFamilyString(t *testing.T) {
	cases := []struct {
		name         string
		family       maroto.Family
		familyString string
	}{
		{
			"maroto.Arial",
			maroto.Arial,
			"arial",
		},
		{
			"maroto.Helvetica",
			maroto.Helvetica,
			"helvetica",
		},
		{
			"maroto.Symbol",
			maroto.Symbol,
			"symbol",
		},
		{
			"maroto.ZapBats",
			maroto.ZapBats,
			"zapfdingbats",
		},
		{
			"maroto.Courier",
			maroto.Courier,
			"courier",
		},
	}

	for _, c := range cases {
		// Act
		familyString := maroto.GetFamilyString(c.family)

		// Assert
		assert.Equal(t, familyString, c.familyString)
	}
}

func TestFont_GetExtensionString(t *testing.T) {
	cases := []struct {
		name            string
		extension       maroto.Extension
		extensionString string
	}{
		{
			"maroto.Jpg",
			maroto.Jpg,
			"jpg",
		},
		{
			"maroto.Png",
			maroto.Png,
			"png",
		},
	}

	for _, c := range cases {
		// Act
		familyString := maroto.GetExtensionString(c.extension)

		// Assert
		assert.Equal(t, familyString, c.extensionString)
	}
}
