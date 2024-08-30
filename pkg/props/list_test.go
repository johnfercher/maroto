package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestList_MakeValid(t *testing.T) {
	t.Run("when MinimumRowsBypage is less than 1, should set MinimumRowsBypage to 1", func(t *testing.T) {
		list := props.List{MinimumRowsBypage: 0}

		list.MakeValid()
		assert.Equal(t, list.MinimumRowsBypage, 1)
	})
}
