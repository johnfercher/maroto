package main

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(false).
		Build()

	m := pkg.NewMaroto(cfg)

	r := row.NewAdaptive(
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(3).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(2).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
	)

	r2 := row.NewAdaptive(col.New(2).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
		WithStyle(&props.Cell{BorderType: border.Full}),
	)

	r3 := row.NewAdaptive(
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(3).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(2).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
	)

	r4 := row.NewAdaptive(
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(3).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(2).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
		col.New(1).Add(text.New("as da s fas df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf df asd fasd fas dfa sdf as d fas dfa sdf asdf")).
			WithStyle(&props.Cell{BorderType: border.Full}),
	)

	m.AddRows(r, r2, r3, r4)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("maroto.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}
