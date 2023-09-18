package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/maroto"
)

func main() {
	m := maroto.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	base64 := document.GetBase64()
	fmt.Println(base64)
}
