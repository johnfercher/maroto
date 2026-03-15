// Package test implements unit test feature.
// nolint:testpackage // that's the integration test of maroto
package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
)

const (
	file = "maroto_test.json"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("when called first, should setup singleton and set t", func(t *testing.T) {
		t.Parallel()
		// Act
		sut := New(t)

		// Assert
		assert.Equal(t, t, sut.t)
	})
	t.Run("when called not first, should use singleton and set t", func(t *testing.T) {
		t.Parallel()
		// Arrange
		_ = New(t)

		// Act
		sut := New(t)

		// Assert
		assert.Equal(t, t, sut.t)
	})
}

func TestMarotoTest_Assert(t *testing.T) {
	t.Parallel()
	t.Run("when call assert, should set node", func(t *testing.T) {
		t.Parallel()
		// Arrange
		n := fixture.Node("maroto")
		sut := New(t)

		// Act
		sut.Assert(n)

		// Assert
		assert.Equal(t, n, sut.node)
	})
}

func TestMarotoTest_Save(t *testing.T) {
	t.Parallel()
	t.Run("when cannot save, should not create file", func(t *testing.T) {
		t.Parallel()
		// Arrange
		file := ""
		n := fixture.Node("maroto")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		// Act
		sut.Equals(file)

		// Assert
		path := configSingleton.getAbsoluteFilePath(file)
		_, err := os.ReadFile(path)
		assert.NotNil(t, err)
		assert.True(t, innerT.Failed())
	})
	t.Run("when can save, should create file", func(t *testing.T) {
		t.Parallel()
		// Arrange
		n := fixture.Node("maroto")
		sut := New(t).Assert(n)

		// Act
		sut.Equals(file)

		// Assert
		path := configSingleton.getAbsoluteFilePath(file)
		bytes, err := os.ReadFile(path)
		assert.Nil(t, err)

		testNode := &Node{}
		_ = json.Unmarshal(bytes, testNode)
		assert.Equal(t, "maroto", testNode.Type)
		assert.Equal(t, "page", testNode.Nodes[0].Type)
	})
}

func TestMarotoTest_Equals(t *testing.T) {
	t.Parallel()
	t.Run("when file saved is not equals to current, should fail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		n := fixture.Node("not_maroto")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		// Act
		sut.Equals(file)

		// Assert
		assert.True(t, innerT.Failed())
	})
	t.Run("when file saved is equals to current, should be success", func(t *testing.T) {
		t.Parallel()
		// Arrange
		n := fixture.Node("maroto")
		innerT := &testing.T{}
		sut := New(innerT).Assert(n)

		// Act
		sut.Equals(file)

		// Assert
		assert.False(t, innerT.Failed())
	})
}
