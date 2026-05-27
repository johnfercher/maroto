// nolint:errchkjson // not needed
package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/johnfercher/go-tree/node"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

var (
	ErrCannotReadDir       = errors.New("cannot read directory")
	ErrCannotReadFile      = errors.New("cannot read file")
	ErrCannotUnmarshallYML = errors.New("cannot unmarshall yaml")
	ErrMarotoYMLNotFound   = errors.New("found go.mod but not .maroto.yml")
)

var (
	marotoFile      = ".maroto.yml"
	goModFile       = "go.mod"
	configSingleton *Config
	configMu        sync.Mutex
)

type Node struct {
	Value   any            `json:"value,omitempty"`
	Type    string         `json:"type"`
	Details map[string]any `json:"details,omitempty"`
	Nodes   []*Node        `json:"nodes,omitempty"`
}

// MarotoTest is the unit test instance.
type MarotoTest struct {
	t    *testing.T
	node *node.Node[core.Structure]
}

// New creates the MarotoTest instance to unit tests.
func New(t *testing.T) *MarotoTest {
	t.Helper()
	loadConfigSingleton(t)

	return &MarotoTest{
		t: t,
	}
}

// loadConfigSingleton initializes the package-level config singleton in a
// goroutine-safe way. Concurrent calls (parallel tests) are serialized by
// configMu so the read/write of configSingleton is race-free.
func loadConfigSingleton(t *testing.T) {
	t.Helper()
	configMu.Lock()
	defer configMu.Unlock()

	if configSingleton != nil {
		return
	}

	path, err := getMarotoConfigFilePath()
	if err != nil {
		assert.Fail(t, "could not find .maroto.yml file. %s"+err.Error())
	}

	cfg, err := loadMarotoConfigFile(path)
	if err != nil {
		assert.Fail(t, "could not parse .maroto.yml. %s"+err.Error())
	}

	cfg.AbsolutePath = path
	configSingleton = cfg
}

// getConfig returns the package-level config singleton under the mutex so
// reads from parallel tests do not race with the write performed during
// initialization.
func getConfig() *Config {
	configMu.Lock()
	defer configMu.Unlock()
	return configSingleton
}

// Assert validates if the structure is the same as defined by Equals method.
func (m *MarotoTest) Assert(structure *node.Node[core.Structure]) *MarotoTest {
	m.node = structure
	return m
}

// Equals defines which file will be loaded to do the comparison.
func (m *MarotoTest) Equals(file string) *MarotoTest {
	m.t.Helper()
	actual := m.buildNode(m.node)
	actualBytes, _ := json.Marshal(actual)
	actualString := string(actualBytes)

	indentedExpectBytes, err := os.ReadFile(getConfig().getAbsoluteFilePath(file))
	if err != nil {
		assert.Fail(m.t, err.Error())
	}

	savedNode := &Node{}
	_ = json.Unmarshal(indentedExpectBytes, savedNode)
	expectedBytes, _ := json.Marshal(savedNode)

	assert.Equal(m.t, string(expectedBytes), actualString)
	return m
}

// Save is an auxiliary method to update the file to be asserted.
func (m *MarotoTest) Save(file string) *MarotoTest {
	actual := m.buildNode(m.node)
	actualBytes, _ := json.MarshalIndent(actual, "", "\t")

	err := os.WriteFile(getConfig().getAbsoluteFilePath(file), actualBytes, os.ModePerm)
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
		return nil, fmt.Errorf("%w: %w", ErrCannotReadFile, err)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotUnmarshallYML, err)
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
		return "", ErrMarotoYMLNotFound
	}

	parentPath := getParentDir(path)
	return getMarotoConfigFilePathRecursive(parentPath)
}

func hasFileInPath(file string, path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("%w: %s", ErrCannotReadDir, err.Error())
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

	var builder strings.Builder
	for _, dir := range dirs {
		builder.WriteString(dir + "/")
	}

	return builder.String()
}
