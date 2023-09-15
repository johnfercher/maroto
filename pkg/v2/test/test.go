package test

import (
	"encoding/json"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/stretchr/testify/assert"
	"testing"
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
