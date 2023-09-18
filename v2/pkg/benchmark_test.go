package pkg_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/text"
	"testing"
)

type testParams struct {
	pages   int
	workers int
}

func buildTestParams() []testParams {
	output := make([]testParams, 0)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 501; j += 50 {
			output = append(output, testParams{
				pages:   j,
				workers: i,
			})
		}
	}
	return output
}

func BenchmarkMarotoGenerateAsync(b *testing.B) {
	for _, param := range buildTestParams() {
		b.Run(fmt.Sprintf("pages=%d,workers=%d", param.pages, param.workers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cfg := config.NewBuilder().WithWorkerPoolSize(param.workers).Build()
				m := pkg.NewMaroto(cfg)
				for i := 0; i < param.pages; i++ {
					m.AddRows(row.New(10).Add(col.New(12).Add(text.New("text"))))
				}
				_, _ = m.Generate()
			}
		})
	}
}
