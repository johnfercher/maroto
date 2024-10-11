package components

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
)

type PdfComponent interface {
	Generate(provider processorprovider.ProcessorProvider) core.Component
}
