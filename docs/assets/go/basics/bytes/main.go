package main

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/maroto"
	"log"
)

func main() {
	m := maroto.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes := document.GetBytes()
	fmt.Println(bytes)
}
