package test

import (
	"encoding/json"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/domain"

	"github.com/johnfercher/go-tree/tree"
	"github.com/stretchr/testify/assert"
)

type Node struct {
	Type  string  `json:"type"`
	Nodes []*Node `json:"nodes,omitempty"`
}

type MarotoTest struct {
	t    *testing.T
	node *tree.Node[domain.Structure]
}

func New(t *testing.T) *MarotoTest {
	return &MarotoTest{
		t: t,
	}
}

func (m *MarotoTest) Assert(maroto domain.Maroto) *MarotoTest {
	m.node = maroto.GetStructure()
	return m
}

func (m *MarotoTest) JSON(expect string) {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.Marshal(actual)

	assert.Equal(m.t, expect, string(actualBytes))
}

func (m *MarotoTest) buildNode(node *tree.Node[domain.Structure]) *Node {
	data := node.GetData()
	actual := &Node{
		Type: data.Type,
	}

	nexts := node.GetNexts()
	for _, next := range nexts {
		actual.Nodes = append(actual.Nodes, m.buildNode(next))
	}

	return actual
}
