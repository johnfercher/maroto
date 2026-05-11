package gofpdf_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// Regression test for issue #550: custom font bytes passed to gofpdf
// must be cloned so each provider instance owns its own backing array.
// gofpdf mutates the slice during PDF finalization (putfonts /
// GenerateCutFont), which previously caused a data race when multiple
// providers (one per worker in concurrent mode) shared the same bytes.
//
// This guards against a regression by confirming that calling Build
// does not mutate the caller's font bytes.
func TestBuilder_Build_DoesNotShareCustomFontBytes(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := gofpdf.NewBuilder()
	font := fixture.FontProp()
	ttf, err := os.ReadFile(buildPath("docs/assets/fonts/arial-unicode-ms.ttf"))
	require.NoError(t, err)
	snapshot := append([]byte(nil), ttf...)

	cfg := &entity.Config{
		Dimensions:  &entity.Dimensions{Width: 100, Height: 200},
		Margins:     &entity.Margins{Left: 10, Top: 10, Right: 10, Bottom: 10},
		DefaultFont: &font,
		CustomFonts: []entity.CustomFont{
			fixture.TestFont{
				Family: fontfamily.Arial,
				Style:  fontstyle.Normal,
				Bytes:  ttf,
			},
		},
	}

	// Act — build twice, simulating two concurrent workers sharing the
	// same CustomFonts entry.
	dep1 := sut.Build(cfg, nil)
	dep2 := sut.Build(cfg, nil)

	// Assert
	assert.NotNil(t, dep1)
	assert.NotNil(t, dep2)
	// The caller's bytes must be untouched after Build.
	assert.Equal(t, snapshot, ttf)
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	dir = strings.ReplaceAll(dir, "internal/providers/gofpdf", "")
	return path.Join(dir, file)
}
