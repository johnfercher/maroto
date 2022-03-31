package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	// m.SetBorder(true)

	byteSlices, err := ioutil.ReadFile("internal/assets/images/biplane.jpg")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	headerSmall, smallContent := getSmallContent()
	headerMedium, mediumContent := getMediumContent()

	base64image := base64.StdEncoding.EncodeToString(byteSlices)

	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				_ = m.Base64Image(base64image, consts.Jpg, props.Rect{
					Center:  true,
					Percent: 70,
				})
			})

			m.ColSpace(3)

			m.Col(3, func() {
				m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
					Center:  true,
					Percent: 75,
				})
			})

			m.Col(3, func() {
				id := "https://github.com/johnfercher/maroto"
				_ = m.Barcode(id, props.Barcode{
					Center:     true,
					Proportion: props.Proportion{Width: 50, Height: 10},
					Percent:    75,
				})
				m.Text(id, props.Text{
					Size:  7,
					Align: consts.Center,
					Top:   14,
				})
			})
		})

		m.Line(1.0,
			props.Line{
				Color: color.Color{
					255, 0, 0,
				},
			})

		m.Row(12, func() {
			m.Col(3, func() {
				_ = m.FileImage("internal/assets/images/gopherbw.png", props.Rect{
					Center: true,
				})
			})

			m.Col(6, func() {
				m.Text("Packages Report: Daily", props.Text{
					Top:   1,
					Align: consts.Center,
				})
				m.Text("Type: Small, Medium", props.Text{
					Top:   7,
					Align: consts.Center,
				})
			})

			m.Col(3, func() {
				m.Text("20/07/1994", props.Text{
					Size:   10,
					Style:  consts.BoldItalic,
					Top:    4,
					Family: consts.Helvetica,
				})
			})
		})

		m.Line(1.0, props.Line{
			Style: consts.Dotted,
			Width: 1.0,
		})

		m.Row(22, func() {
			m.Col(0, func() {
				m.Text(fmt.Sprintf("Small: %d, Medium %d", len(smallContent), len(mediumContent)), props.Text{
					Size:  15,
					Style: consts.Bold,
					Align: consts.Center,
					Top:   4,
					Color: color.Color{
						Blue: 180,
					},
				})
				m.Text("Brasil / São Paulo", props.Text{
					Size:  12,
					Align: consts.Center,
					Top:   12,
				})
			})
		})

		m.Line(1.0, props.Line{
			Style: consts.Dashed,
			Width: 0.5,
		})
	})

	m.RegisterFooter(func() {
		m.Row(40, func() {
			m.Col(4, func() {
				m.Signature("Signature 1", props.Font{
					Family: consts.Courier,
					Style:  consts.BoldItalic,
					Size:   9,
					Color: color.Color{
						Red: 200,
					},
				})
			})

			m.Col(4, func() {
				m.Signature("Signature 2")
			})

			m.Col(4, func() {
				m.Signature("Signature 3")
			})
		})
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(strconv.Itoa(m.GetCurrentPage())+"/{nb}", props.Text{
					Align: consts.Right,
					Size:  8,
				})
			})
		})
	})

	m.Row(15, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Small Packages / %du.", len(smallContent)), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerSmall, smallContent, props.TableList{
		ContentProp: props.TableListContent{
			GridSizes: []uint{3, 6, 3},
			Color:     color.Color{100, 0, 0},
		},
		HeaderProp: props.TableListContent{
			GridSizes: []uint{3, 6, 3},
		},
		AlternatedBackground: &color.Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		},
	})

	m.AddPage()

	m.Row(15, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Medium Packages / %du.", len(mediumContent)), props.Text{
				Top:   8,
				Style: consts.Bold,
			})
		})
	})

	m.TableList(headerMedium, mediumContent, props.TableList{
		ContentProp: props.TableListContent{
			Family:    consts.Courier,
			Style:     consts.Italic,
			GridSizes: []uint{5, 5, 2},
		},
		HeaderProp: props.TableListContent{
			GridSizes: []uint{5, 5, 2},
			Family:    consts.Courier,
			Style:     consts.BoldItalic,
			Color:     color.Color{100, 0, 0},
		},
		Align: consts.Center,
		Line:  true,
	})

	err = m.OutputFileAndClose("internal/examples/pdfs/sample1.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getSmallContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "Cost"}

	contents := [][]string{}
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 20,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 25,00"})
	contents = append(contents, []string{"São José do Vale do Rio Preto", "Osasco", "R$ 20,00"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 5,00"})
	contents = append(contents, []string{"Congonhas", "Fortaleza", "R$ 100,00"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 200,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 44,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 56,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 35,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 82,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 62,00"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 21,00"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 12,00"})
	contents = append(contents, []string{"Niterói", "Itapevi", "R$ 21,00"})
	contents = append(contents, []string{"São Paulo", "Rio de Janeiro", "R$ 31,00"})
	contents = append(contents, []string{"São Carlos", "Petrópolis", "R$ 42,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 19,00"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 42,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 58,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 21,00"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 56,00"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 35,00"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 82,00"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 62,00"})

	return header, contents
}

func getMediumContent() ([]string, [][]string) {
	header := []string{"Origin", "Destiny", "Cost per Hour"}

	contents := [][]string{}
	contents = append(contents, []string{"São José do Vale do Rio Preto", "Osasco", "R$ 12,00"})
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
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,30"})
	contents = append(contents, []string{"Natal", "Jundiaí", "R$ 1,10"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})
	contents = append(contents, []string{"Campinas", "Recife", "R$ 5,80"})
	contents = append(contents, []string{"São Vicente", "Juiz de Fora", "R$ 3,90"})
	contents = append(contents, []string{"Taubaté", "Rio de Janeiro", "R$ 7,70"})
	contents = append(contents, []string{"Suzano", "Petrópolis", "R$ 6,40"})
	contents = append(contents, []string{"Jundiaí", "Florianópolis", "R$ 2,00"})
	contents = append(contents, []string{"Florianópolis", "Osasco", "R$ 1,90"})
	contents = append(contents, []string{"Osasco", "São Paulo", "R$ 0,70"})
	contents = append(contents, []string{"Congonhas", "São José do Vale do Rio Preto", "R$ 11,30"})
	contents = append(contents, []string{"Natal", "Santo André", "R$ 19,80"})
	contents = append(contents, []string{"Rio Grande do Norte", "Sorocaba", "R$ 4,20"})

	return header, contents
}
