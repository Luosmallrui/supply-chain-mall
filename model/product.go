package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Price         float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	OriginalPrice float64        `gorm:"type:decimal(10,2);not null" json:"originalPrice"`
	Image         string         `gorm:"type:varchar(512)" json:"image"`
	Brand         string         `gorm:"type:varchar(100)" json:"brand"`
	BrandSubtitle string         `gorm:"type:varchar(255)" json:"brandSubtitle"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Product) TableName() string {
	return "product"
}
