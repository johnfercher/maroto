package loader_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/loader"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Run("when invalid extension sent, should return ErrUnsupportedExtension", func(t *testing.T) {
		_, err := loader.NewLoader().Load("README.md")
		assert.ErrorIs(t, err, loader.ErrUnsupportedExtension)
	})

	t.Run("when invalid path sent, should return ErrInvalidPath", func(t *testing.T) {
		_, err := loader.NewLoader().Load("http://hi this is an invalid path.png")
		assert.ErrorIs(t, err, loader.ErrInvalidPath)
	})

	t.Run("when path with unsupported protocol sent, should return ErrSupportedProtocol", func(t *testing.T) {
		_, err := loader.NewLoader().Load("irc://foobar.com/asset.png")
		assert.ErrorIs(t, err, loader.ErrUnsupportedProtocol)
	})

	t.Run("when valid local path sent, should return bytes of file", func(t *testing.T) {
		p, err := loader.NewLoader().Load("../../../docs/assets/images/logo.png")
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})

	t.Run("when valid file uri sent, should return bytes of file", func(t *testing.T) {
		p, err := loader.NewLoader().Load("file://../../../docs/assets/images/logo.png")
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})

	t.Run("when valid network path sent, should return bytes of asset", func(t *testing.T) {
		p, err := loader.NewLoader().Load("https://www.iana.org/_img/2013.1/rir-map.svg")
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})
}
