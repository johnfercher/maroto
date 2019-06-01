package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMaroto(t *testing.T) {
	m := maroto.NewMaroto(maroto.Vertical, maroto.A4)

	assert.NotNil(t, m)
	assert.Equal(t, fmt.Sprintf("%T", m), "*maroto.maroto")
}
