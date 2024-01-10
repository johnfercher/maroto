// Package implements unit test feature.
package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/stretchr/testify/assert"
)

var (
	marotoFile              = ".maroto.yml"
	goModFile               = "go.mod"
	configSingleton *Config = nil
)

type Node struct {
	Value   interface{}            `json:"value,omitempty"`
	Type    string                 `json:"type"`
	Details map[string]interface{} `json:"details,omitempty"`
	Nodes   []*Node                `json:"nodes,omitempty"`
}

type MarotoTest struct {
	t    *testing.T
	node *node.Node[core.Structure]
}

func New(t *testing.T) *MarotoTest {
	if configSingleton == nil {
		path, err := getMarotoConfigFilePath()
		if err != nil {
			assert.Fail(t, fmt.Sprintf("could not find .maroto.yml file. %s", err.Error()))
		}

		cfg, err := loadMarotoConfigFile(path)
		if err != nil {
			assert.Fail(t, fmt.Sprintf("could not parse .maroto.yml. %s", err.Error()))
		}

		cfg.AbsolutePath = path
		configSingleton = cfg
	}

	return &MarotoTest{
		t: t,
	}
}

func (m *MarotoTest) Assert(structure *node.Node[core.Structure]) *MarotoTest {
	m.node = structure
	return m
}

func (m *MarotoTest) Equals(file string) *MarotoTest {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.Marshal(actual)
	actualString := string(actualBytes)

	indentedExpectBytes, err := os.ReadFile(configSingleton.getAbsoluteFilePath(file))
	if err != nil {
		assert.Fail(m.t, err.Error())
	}

	savedNode := &Node{}
	_ = json.Unmarshal(indentedExpectBytes, savedNode)
	expectedBytes, _ := json.Marshal(savedNode)

	assert.Equal(m.t, string(expectedBytes), actualString)
	return m
}

func (m *MarotoTest) Save(file string) *MarotoTest {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.MarshalIndent(actual, "", "\t")

	err := os.WriteFile(configSingleton.getAbsoluteFilePath(file), actualBytes, os.ModePerm)
	if err != nil {
		assert.Fail(m.t, err.Error())
	}

	return m
}

func (m *MarotoTest) buildNode(node *node.Node[core.Structure]) *Node {
	data := node.GetData()
	actual := &Node{
		Type:    data.Type,
		Value:   data.Value,
		Details: data.Details,
	}

	nexts := node.GetNexts()
	for _, next := range nexts {
		actual.Nodes = append(actual.Nodes, m.buildNode(next))
	}

	return actual
}

func getMarotoConfigFilePath() (string, error) {
	path, _ := os.Getwd()
	path += "/"

	return getMarotoConfigFilePathRecursive(path)
}

func loadMarotoConfigFile(path string) (*Config, error) {
	bytes, err := os.ReadFile(path + "/" + marotoFile)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func getMarotoConfigFilePathRecursive(path string) (string, error) {
	hasMaroto, err := hasFileInPath(marotoFile, path)
	if err != nil {
		return "", err
	}

	if hasMaroto {
		return path, nil
	}

	hasGoMod, err := hasFileInPath(goModFile, path)
	if err != nil {
		return "", err
	}

	if hasGoMod {
		return "", errors.New("found go.mod but not .maroto.yml")
	}

	parentPath := getParentDir(path)
	return getMarotoConfigFilePathRecursive(parentPath)
}

func hasFileInPath(file string, path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		if entry.Name() == file {
			return true, nil
		}
	}

	return false, nil
}

func getParentDir(path string) string {
	dirs := strings.Split(path, "/")
	dirs = dirs[:len(dirs)-2]

	var newPath string
	for _, dir := range dirs {
		newPath += dir + "/"
	}

	return newPath
}
