package dao

import (
	"gorm.io/gorm"
	"supply-chain-mall/types"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// Create 创建商品
func (r *ProductRepo) Create(user *types.User) error {
	return r.db.Create(user).Error
}

// Update 更新商品
func (r *ProductRepo) Update(user *types.User) error {
	return r.db.Save(user).Error
}

// FindByID 通过ID查找场频
func (r *ProductRepo) FindByID(id uint) (*types.User, error) {
	var user types.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
