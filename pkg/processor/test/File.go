package test

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	marotoFile          = ".maroto.yml"
	goModFile           = "go.mod"
	fileSingleton *File = nil
)

type File struct {
	config *Config
}

func NewFileReader() *File {
	if fileSingleton == nil {
		path, err := getMarotoConfigFilePath()
		if err != nil {
			return nil
		}

		cfg, err := loadMarotoConfigFile(path)
		if err != nil {
			return nil
		}

		cfg.AbsolutePath = path
		fileSingleton = &File{config: cfg}
	}

	return fileSingleton
}

func (f File) LoadFile(file string) ([]byte, error) {
	fileContent, err := os.ReadFile(f.config.getAbsoluteFilePath(file))
	if err != nil {
		return nil, err
	}
	return fileContent, err
}

func getMarotoConfigFilePath() (string, error) {
	path, _ := os.Getwd()
	path += "/"

	return getMarotoConfigFilePathRecursive(path)
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
