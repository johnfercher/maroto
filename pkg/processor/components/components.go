package components

import (
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/provider"
)

type Component interface {
	Generate(provider provider.Provider) core.Component
}
