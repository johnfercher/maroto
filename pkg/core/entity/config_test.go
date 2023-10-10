package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_ToMap(t *testing.T) {
	// Arrange
	sut := fixtureConfig()

	// Act
	m := sut.ToMap()

	// Assert
	assert.Equal(t, provider.Gofpdf, m["config_provider_type"])
	assert.Equal(t, 100.0, m["maroto_dimension_width"])
	assert.Equal(t, 200.0, m["maroto_dimension_height"])
	assert.Equal(t, 20.0, m["config_margin_left"])
	assert.Equal(t, 30.0, m["config_margin_top"])
	assert.Equal(t, 40.0, m["config_margin_right"])
	assert.Equal(t, 50.0, m["config_margin_bottom"])
	assert.Equal(t, fontfamily.Helvetica, m["prop_font_family"])
	assert.Equal(t, fontstyle.Bold, m["prop_font_style"])
	assert.Equal(t, 15.0, m["prop_font_size"])
	assert.Equal(t, "RGB(255, 0, 0)", m["prop_font_color"])
	assert.Equal(t, 7, m["config_workers"])
	assert.Equal(t, true, m["config_debug"])
	assert.Equal(t, 15, m["config_max_grid_sum"])
	assert.Equal(t, "pattern", m["config_page_number_pattern"])
	assert.Equal(t, props.South, m["config_page_number_place"])
	assert.Equal(t, protection.Print, m["config_protection_type"])
	assert.Equal(t, "654321", m["config_user_password"])
	assert.Equal(t, "123456", m["config_owner_password"])
	assert.Equal(t, true, m["config_compression"])
	assert.Equal(t, "Utf8Text(author, true)", m["config_metadata_author"])
	assert.Equal(t, "Utf8Text(creator, false)", m["config_metadata_creator"])
	assert.Equal(t, "Utf8Text(subject, true)", m["config_metadata_subject"])
	assert.Equal(t, "Utf8Text(title, true)", m["config_metadata_title"])
	assert.Equal(t, true, m["config_metadata_creation_date"])
	assert.Equal(t, "[1 2 3]", m["entity_image_bytes"])
	assert.Equal(t, extension.Png, m["entity_extension"])
	assert.Equal(t, 100.0, m["background_dimension_width"])
	assert.Equal(t, 200.0, m["background_dimension_height"])
}

func fixtureConfig() Config {
	dimensions := fixtureDimensions()
	margins := fixtureMargins()
	font := fixtureFont()
	protection := fixtureProtection()
	metadata := fixtureMetadata()
	image := fixtureImage()

	return Config{
		ProviderType:      provider.Gofpdf,
		Dimensions:        &dimensions,
		Margins:           &margins,
		DefaultFont:       &font,
		Workers:           7,
		Debug:             true,
		MaxGridSize:       15,
		PageNumberPattern: "pattern",
		PageNumberPlace:   props.South,
		Protection:        &protection,
		Compression:       true,
		Metadata:          &metadata,
		BackgroundImage:   &image,
	}
}

func fixtureFont() props.Font {
	return props.Font{
		Family: fontfamily.Helvetica,
		Style:  fontstyle.Bold,
		Size:   15,
		Color:  &props.RedColor,
	}
}
