package providers

import (
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type ProviderOption func(p core.Provider)

func WithCache(cache cache.Cache) ProviderOption {
	return func(p core.Provider) {
		p.SetCache(cache)
	}
}
