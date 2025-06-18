package svc

import (
	"context"
	"supply-chain-mall/dao"
	"supply-chain-mall/model"
	"supply-chain-mall/pkg/cache"
	"supply-chain-mall/types"
)

var _ IProductService = (*ProductService)(nil)

type ProductService struct {
	ProductRepo *dao.ProductRepo
	Cache       *cache.TaskCache
}

type IProductService interface {
	CreateProduct(ctx context.Context, p *types.Product) error
	FindAllProduct(ctx context.Context) ([]*model.Product, error)
}

func (s *ProductService) CreateProduct(ctx context.Context, p *types.Product) error {
	err := s.ProductRepo.Create(ctx, p)
	return err
}

func (s *ProductService) FindAllProduct(ctx context.Context) ([]*model.Product, error) {
	products, err := s.ProductRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}
