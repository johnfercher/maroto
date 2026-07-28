// nolint:misspell // arabic is being classified as english misspells
package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/checkbox"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/fontrepository"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto("docs/assets/fonts/arial-unicode-ms.ttf")
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/arabicv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/arabicv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto(customFontFile string) core.Maroto {
	// Arabic needs a font that carries the Arabic Presentation Forms-B glyphs,
	// which the standard PDF fonts do not have.
	customFont := "arial-unicode-ms"

	customFonts, err := fontrepository.New().
		AddUTF8Font(customFont, fontstyle.Normal, customFontFile).
		AddUTF8Font(customFont, fontstyle.Bold, customFontFile).
		Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := config.NewBuilder().
		WithCustomFonts(customFonts).
		WithDefaultFont(&props.Font{Family: customFont}).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(text.NewRow(10, "Arabic without RTL, and with RTL", props.Text{
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	// The same string twice: once as maroto renders it today, once with the
	// shaping and the reordering turned on.
	const greeting = "مرحبا بالعالم"

	m.AddRow(10,
		text.NewCol(6, greeting, props.Text{Align: align.Center}),
		text.NewCol(6, greeting, props.Text{Align: align.Center, RTL: true}),
	)

	// Right alignment is the natural one for Arabic.
	m.AddRows(text.NewRow(10, "لغة عربية", props.Text{
		Align: align.Right,
		RTL:   true,
	}))

	// Arabic mixed with Latin words and numbers: each run keeps its own
	// direction, so "Maroto" and "42" stay readable inside the Arabic.
	m.AddRows(text.NewRow(10, "المكتبة Maroto تدعم 42 لغة", props.Text{
		Align: align.Right,
		RTL:   true,
	}))

	// A paragraph long enough to wrap. Lines are broken first and each one is
	// shaped and reordered on its own, so every line reads correctly.
	longText := "اللغة العربية هي أكثر اللغات السامية تحدثا وإحدى أكثر اللغات انتشارا في العالم " +
		"يتحدثها أكثر من أربعمائة مليون نسمة ويتوزع متحدثوها في المنطقة المعروفة باسم الوطن العربي"

	m.AddRow(40,
		text.NewCol(12, longText, props.Text{Align: align.Right, RTL: true}),
	)

	m.AddRows(checkbox.NewRow(8, "أوافق على الشروط", props.Checkbox{
		Checked: true,
		Size:    5,
		RTL:     true,
	}))

	return m
}
