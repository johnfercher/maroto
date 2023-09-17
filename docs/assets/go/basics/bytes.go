package main

import (
	"fmt"
	"log"

	v2 "github.com/johnfercher/maroto/pkg/v2"
)

func main() {
	m := v2.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes := document.GetBytes()
	fmt.Println(bytes)
}
