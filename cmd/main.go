package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"log"
)

func main() {
	m := v2.NewMaroto("v2.pdf")

	// Add things

	err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
func buildMarotoPDF() domain.MarotoMetrified {
	m := v2.NewMaroto("v2.pdf")
	return v2.NewMarotoMetrified(m)
}

func buildMarotoHTML() domain.MarotoMetrified {
	builder := config.NewBuilder().
		WithPageSize(config.A4).
		WithProvider(provider.HTML)

	m := v2.NewMaroto("v2.html", builder)
	return v2.NewMarotoMetrified(m)
}

func gen(m domain.MarotoMetrified) {
	for _ = range [10]int{} {
		m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	}
	//m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	//m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())

	report.txt, err := m.GenerateWithReport()
	if err != nil {
		log.Fatal(err.Error())
	}

	report.txt.Print()
}

func buildCodesRow() domain.Row {
	r := row.New(70)

	col1 := col.New(4)
	col1.Add(code.NewBar("barcode"))

	col2 := col.New(4)
	col2.Add(code.NewQr("qrcode"))

	col3 := col.New(4)
	col3.Add(code.NewMatrix("matrixcode"))

	r.Add(col1, col2, col3)
	return r
}

func buildImagesRow() domain.Row {
	row := row.New(70)

	col1 := col.New(6)
	col1.Add(image.NewFromFile("internal/assets/images/frontpage.png"))

	byteSlices, err := os.ReadFile("internal/assets/images/frontpage.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)
	col2 := col.New(6)
	col2.Add(image.NewFromBase64(stringBase64, consts.Png))

	row.Add(col1, col2)

	return row
}

func buildTextsRow() domain.Row {
	row := row.New(70)

	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	col1 := col.New(6)
	col1.Add(text.New(colText, props.Text{
		Align: consts.Center,
	}))

	col2 := col.New(6)
	col2.Add(signature.New("Fulano de Tal", props.Font{
		Style:  consts.Italic,
		Size:   20,
		Family: consts.Courier,
	}))

	row.Add(col1, col2)

	return row
}
*/
