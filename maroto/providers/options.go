package providers

import (
	"github.com/johnfercher/v2/maroto/cache"
	"github.com/johnfercher/v2/maroto/domain"
)

type ProviderOption func(p domain.Provider)

func WithCache(cache cache.Cache) ProviderOption {
	return func(p domain.Provider) {
		p.SetCache(cache)
	}
}
