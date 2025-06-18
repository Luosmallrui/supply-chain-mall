//go:build wireinject

package main

import (
	"github.com/google/wire"
	"supply-chain-mall/controller"
	"supply-chain-mall/pkg/core"
)

func NewInjector() (*core.AppProvider, error) {
	panic(
		wire.Build(
			controller.ProviderSet,
			core.ProviderSet,
		),
	)
}
