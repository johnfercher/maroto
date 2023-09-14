package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/providers"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
)

func main() {
	provider := providers.NewGofpdf()
	pdf := v2.NewDocument(provider, "v2.pdf")

	provider = providers.NewHTML()
	pdf = v2.NewDocument(provider, "v2.html")

	//pdf.Add(buildCodesRow())
	//pdf.Add(buildImageRow())
	pdf.Add(buildTextRow(), buildTextRow(), buildTextRow(), buildTextRow(), buildTextRow())

	//pdf.Add()
	//pdf.ForceAddPage(p)

	err := pdf.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*func buildCodesRow() domain.Row {
	r := row.New(70)

	//image := image.New("image1")

	col1 := col.New(4)
	col1.Add(barcode.New("barcode"))

	col2 := col.New(4)
	col2.Add(qrcode.New("qrcode"))

	col3 := col.New(4)
	col3.Add(matrixcode.New("matrixcode"))

	r.Add(col1, col2, col3)
	return r
}

func buildImageRow() domain.Row {
	row := row.New(70)

	col1 := col.New(6)
	col1.Add(image.NewFromFile("internal/assets/images/biplane.jpg"))

	byteSlices, err := os.ReadFile("internal/assets/images/gopherbw.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)
	col2 := col.New(6)
	col2.Add(image.NewFromBase64(stringBase64, consts.Png))

	row.Add(col1, col2)

	return row
}*/

func buildTextRow() domain.Row {
	row := row.New(70)

	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	col1 := col.New(6)
	col1.Add(text.New(colText, props.Text{
		Align: consts.Center,
	}))

	/*col2 := col.New(6)
	col2.Add(signature.New("Fulano de Tal", props.Font{
		Style:  consts.Italic,
		Size:   20,
		Family: consts.Courier,
	}))*/

	row.Add(col1, col1)

	return row
}
