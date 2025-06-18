package svc

import (
	"supply-chain-mall/dao"
	"supply-chain-mall/pkg/cache"
	"supply-chain-mall/types"
	"time"
)

var _ IProductService = (*ProductService)(nil)

type ProductService struct {
	UserRepo *dao.ProductRepo
	Cache    *cache.TaskCache
}

type IProductService interface {
	CreateProduct() error
}

func (s *ProductService) CreateProduct() error {
	s.UserRepo.Create(&types.User{
		ID:        0,
		Username:  "",
		Nickname:  "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	return nil
}
