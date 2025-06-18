package dao

import (
	"context"
	"gorm.io/gorm"
	"supply-chain-mall/model"
	"supply-chain-mall/types"
)

type ProductRepo struct {
	Repo[model.Product]
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{Repo: NewRepo[model.Product](db)}
}

// Create 创建商品
func (r *ProductRepo) Create(ctx context.Context, p *types.Product) error {
	return r.Repo.Create(ctx, &model.Product{
		Name:          p.Name,
		Price:         p.Price,
		OriginalPrice: float64(p.OriginalPrice),
		Image:         p.Image,
		Brand:         p.Brand,
		BrandSubtitle: p.BrandSubtitle,
	})
}

func (r *ProductRepo) FindAll(ctx context.Context) ([]*model.Product, error) {
	return r.Repo.FindAll(ctx, func(db *gorm.DB) {
		db.Where(&model.Product{})
	})
}
