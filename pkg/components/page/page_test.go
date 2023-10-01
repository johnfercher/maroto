package page_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// Act
	sut := page.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*page.page", fmt.Sprintf("%T", sut))
}
