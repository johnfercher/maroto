package main

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg"
	"log"
)

func main() {
	m := pkg.NewMaroto()

	// Do things

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	base64 := document.GetBase64()
	fmt.Println(base64)
}
