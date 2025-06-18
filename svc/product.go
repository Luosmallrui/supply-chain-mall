package svc

import (
	"context"
	"supply-chain-mall/dao"
	"supply-chain-mall/model"
	"supply-chain-mall/pkg/cache"
)

var _ IProductService = (*ProductService)(nil)

type ProductService struct {
	ProductRepo *dao.ProductRepo
	Cache       *cache.TaskCache
}

type IProductService interface {
	CreateProduct() error
	FindAllProduct(ctx context.Context) ([]*model.Product, error)
}

func (s *ProductService) CreateProduct() error {
	return nil
}

func (s *ProductService) FindAllProduct(ctx context.Context) ([]*model.Product, error) {
	products, err := s.ProductRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, err
}
