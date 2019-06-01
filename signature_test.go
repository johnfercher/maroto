package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSignature(t *testing.T) {
	signature := maroto.NewSignature(&mocks.Pdf{}, &mocks.Math{}, &mocks.Text{})

	assert.NotNil(t, signature)
	assert.Equal(t, fmt.Sprintf("%T", signature), "*maroto.signature")
}
