package controller

import (
	"github.com/google/wire"
	"supply-chain-mall/svc"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(Product), "*"),
	svc.ProviderSet,
	wire.Struct(new(Controllers), "*"),
	NewGinServer,
)
