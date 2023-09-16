package v2_test

import (
	"fmt"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/text"
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
				maroto := v2.NewMaroto(cfg)
				for i := 0; i < param.pages; i++ {
					maroto.AddRows(row.New(10).Add(col.New(12).Add(text.New("text"))))
				}
				maroto.Generate()
			}
		})
	}
}
