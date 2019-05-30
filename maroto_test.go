package maroto_test

import (
	"github.com/johnfercher/maroto"
	"testing"
)

func BenchmarkMaroto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pdf()
	}
}

func Pdf() {
	m := maroto.NewMaroto(maroto.Vertical, maroto.A4)
	header, contents := getContents()

	m.Row("MeliBarcode", 20, func() {
		m.Col("Logo", func() {
			m.Image("assets/images/mercado_livre.png", 4)
		})

		m.ColSpaces(2)

		m.Col("Barcode", func() {
			id := "123456789"
			_ = m.Barcode(id, 30, 9, 5)
			m.Text(id, maroto.Arial, maroto.Bold, 8, 17, maroto.CenterH)
		})
	})

	m.Line()

	m.Row("Destiny", 12, func() {
		m.Col("Logo", func() {
			m.Image("assets/images/mercado_livre.png", 1)
		})

		m.ColSpace()

		m.Col("Packages", func() {
			m.Text("Vendedor: The Collector", maroto.Arial, maroto.Normal, 9, 5, maroto.Left)
			m.Text("Endereco: Nowhere", maroto.Arial, maroto.Normal, 9, 9, maroto.Left)
		})

		m.ColSpace()

		m.Col("Route", func() {
			m.Text("ROUTE.XDA.6", maroto.Arial, maroto.Bold, 15, 7.5, maroto.Left)
		})
	})

	m.Line()

	m.Row("Packages Title", 22, func() {
		m.ColSpaces(2)

		m.Col("Packages", func() {
			m.Text("24", maroto.Arial, maroto.Bold, 20, 10.5, maroto.CenterH)
			m.Text("Pacotes Devolvidos", maroto.Arial, maroto.Normal, 12, 16, maroto.CenterH)
		})

		m.ColSpaces(2)
	})

	m.Line()

	m.RowTableList("Packages", header, contents)

	m.Row("Signature", 15, func() {
		m.Col("Carrier", func() {
			m.Sign("Transportadora", maroto.Arial, maroto.Bold, 8)
		})

		m.ColSpace()

		m.Col("LogisticOperator", func() {
			m.Sign("Operador Logistico", maroto.Arial, maroto.Bold, 8)
		})

		m.ColSpace()

		m.Col("Seller", func() {
			m.Sign("Vendedor", maroto.Arial, maroto.Bold, 8)
		})
	})

	m.OutputFileAndClose("pdf")
}

func getContents() ([]string, [][]string) {
	header := []string{"Envio", "Venda", "Comprador", "Motivo"}

	contents := [][]string{
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
		{"678445", "678543", "Thanos", "Produto queimado"},
		{"489423", "579894", "Peter Parker", "Compra cancelada"},
		{"679076", "272747", "Thor", "Produto errado"},
		{"854364", "996634", "Nebula", "Fraude"},
		{"679095", "768690", "Steve Rogers", "Venda cancelada"},
		{"234512", "356469", "Tony Stark", "Produto errado"},
		{"123451", "996755", "Steve Strange", "Produto errado"},
		{"675523", "352364", "Star Lord", "Compra cancelada"},
		{"787894", "693595", "Gamora", "Fraude"},
		{"908907", "967867", "Scott Lang", "Compra cancelada"},
		{"876453", "797934", "Hank Pyn", "Produto errado"},
	}

	return header, contents
}
