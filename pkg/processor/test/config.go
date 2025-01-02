// Package test implements unit test feature.
package test

// Config is the representation of a test config.
type Config struct {
	AbsolutePath string
	TestPath     string `yaml:"test_path"`
}

func (c *Config) getAbsoluteFilePath(file string) string {
	return c.AbsolutePath + c.TestPath + file
}
