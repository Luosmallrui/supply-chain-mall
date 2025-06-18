package dao

import (
	"github.com/google/wire"
	"supply-chain-mall/pkg/cache"
)

var ProviderSet = wire.NewSet(
	NewProductRepo,
	cache.NewTaskCache,
)
