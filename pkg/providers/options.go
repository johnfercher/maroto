package providers

import (
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/domain"
)

type ProviderOption func(p domain.Provider)

func WithCache(cache cache.Cache) ProviderOption {
	return func(p domain.Provider) {
		p.SetCache(cache)
	}
}
