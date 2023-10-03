package test

import (
	"fmt"
	"log"
	"testing"
)

func TestMarotoTest_GetMarotoConfigFilePath(t *testing.T) {
	// Act
	path, err := getMarotoConfigFilePath()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := loadMarotoConfigFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg)

	fmt.Println(path)
}
