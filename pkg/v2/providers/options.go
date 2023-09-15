package providers

import (
	"github.com/johnfercher/maroto/pkg/v2/cache"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type ProviderOption func(p domain.Provider)

func WithCache(cache cache.Cache) ProviderOption {
	return func(p domain.Provider) {
		p.SetCache(cache)
	}
}
