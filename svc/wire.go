package svc

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(ProductService), "*"),
	wire.Bind(new(IProductService), new(*ProductService)),
)
