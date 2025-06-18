//go:build wireinject

package main

import (
	"github.com/google/wire"
	"supply-chain-mall/config"
	"supply-chain-mall/controller"
	"supply-chain-mall/dao"
	"supply-chain-mall/pkg/client"
	"supply-chain-mall/pkg/core"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet,
	controller.ProviderSet,
	core.ProviderSet,
	client.ProviderSet,
	dao.ProviderSet,
)

func NewInjector() (*core.AppProvider, error) {
	panic(wire.Build(ProviderSet))
}
