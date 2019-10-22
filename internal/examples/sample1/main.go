package main

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"io/ioutil"
	"os"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	byteSlices, err := ioutil.ReadFile("internal/assets/images/biplane.jpg")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	base64 := base64.StdEncoding.EncodeToString(byteSlices)

	headerSmall, smallContent := getSmallContent()
	headerMedium, mediumContent := getMediumContent()

	m.RegisterHeader(func() {

		m.Row(20, func() {
			m.Col(func() {
				m.Base64Image(base64, consts.Jpg, props.Rect{
					Center:  true,
					Percent: 70,
				})
			})

			m.ColSpaces(2)

			m.Col(func() {
				m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
					Percent: 75,
				})
			})

			m.Col(func() {
				id := "https://github.com/johnfercher/maroto"
				_ = m.Barcode(id, props.Barcode{
					Proportion: props.Proportion{Width: 50, Height: 10},
					Percent:    75,
				})
				m.Text(id, props.Text{
					Size:  7,
					Align: consts.Center,
					Top:   16,
				})
			})
		})

		m.Line(1.0)

		m.Row(12, func() {
			m.Col(func() {
				_ = m.FileImage("internal/assets/images/goherbw.png", props.Rect{
					Center: true,
				})
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("Packages Report: Daily", props.Text{
					Top: 4,
				})
				m.Text("Type: Small, Medium", props.Text{
					Top: 10,
				})
			})

			m.ColSpace()

			m.Col(func() {
				m.Text("20/07/1994", props.Text{
					Size:   10,
					Style:  consts.BoldItalic,
					Top:    7.5,
					Family: consts.Helvetica,
				})
			})
		})

		m.Line(1.0)

		m.Row(22, func() {
			m.Col(func() {
				m.Text(fmt.Sprintf("Small: %d, Medium %d", len(smallContent), len(mediumContent)), props.Text{
					Size:  15,
					Style: consts.Bold,
					Align: consts.Center,
					Top:   9,
				})
				m.Text("Brasil / São Paulo", props.Text{
					Size:  12,
					Align: consts.Center,
					Top:   17,
				})
			})
		})

		m.Line(1.0)

	})

	m.RegisterFooter(func() {
		m.Row(40, func() {
			m.Col(func() {
				m.Signature("Signature 1", props.Font{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9,
				})
			})

			m.Col(func() {
				m.Signature("Signature 2")
			})

			m.Col(func() {
				m.Signature("Signature 3")
			})
		})
	})

	m.Row(15, func() {
		m.Col(func() {
			m.Text("Small Packages / 39u.", props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerSmall, smallContent)

	m.Row(15, func() {
		m.Col(func() {
			m.Text("Medium Packages / 22u.", props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerMedium, mediumContent, props.TableList{
		Align: consts.Center,
		HeaderProp: props.Font{
			Family: consts.Courier,
			Style:  consts.BoldItalic,
		},
		ContentProp: props.Font{
			Family: consts.Courier,
			Style:  consts.Italic,
		},
	})

	err = m.OutputFileAndClose("internal/examples/pdfs/sample1.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
}

func getSmallContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "", "Cost"}

	contents := [][]string{}
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "", "R$ 20,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "", "R$ 25,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "", "R$ 20,00"})
	contents = append(contents, []string{"Osasco", "São Paulo", "", "R$ 5,00"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "", "R$ 100,00"})
	contents = append(contents, []string{"Natal", "Santo André", "", "R$ 200,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "", "R$ 44,00"})
	contents = append(contents, []string{"Campinas", "Recife", "", "R$ 56,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "", "R$ 35,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "", "R$ 82,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "", "R$ 62,00"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "", "R$ 21,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "", "R$ 12,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "", "R$ 21,00"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "", "R$ 31,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "", "R$ 42,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "", "R$ 19,00"})
	contents = append(contents, []string{"Osasco", "São Paulo", "", "R$ 7,00"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "", "R$ 113,00"})
	contents = append(contents, []string{"Natal", "Santo André", "", "R$ 198,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "", "R$ 42,00"})
	contents = append(contents, []string{"Campinas", "Recife", "", "R$ 58,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "", "R$ 39,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "", "R$ 77,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "", "R$ 64,00"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "", "R$ 20,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "", "R$ 18,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "", "R$ 24,00"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "", "R$ 23,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "", "R$ 11,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "", "R$ 28,00"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "", "R$ 19,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "", "R$ 23,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "", "R$ 21,00"})
	contents = append(contents, []string{"Osasco", "São Paulo", "", "R$ 6,00"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "", "R$ 109,00"})
	contents = append(contents, []string{"Natal", "Santo André", "", "R$ 244,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "", "R$ 34,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "", "R$ 21,00"})

	return header, contents
}

func getMediumContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "Cost per Hour"}

	contents := [][]string{}
	contents = append(contents, []string{"Niterói", "Itapevi", "R$ 2,10"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 3,10"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 4,20"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 1,90"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 0,70"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "R$ 11,30"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 5,80"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 3,90"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 7,70"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 6,40"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 1,80"})
	contents = append(contents, []string{"Niterói", "Itapevi", "R$ 2,40"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,30"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 1,10"})

	return header, contents
}
