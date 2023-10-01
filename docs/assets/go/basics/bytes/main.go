package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"
)

func main() {
	m := maroto.New()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes := document.GetBytes()
	fmt.Println(bytes)
}
