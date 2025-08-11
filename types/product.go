package types

import "time"

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`         // 改为浮点数
	OriginalPrice int     `json:"originalPrice"` // 改为整数
	Image         string  `json:"image"`
	Brand         string  `json:"brand"`
	BrandSubtitle string  `json:"brandSubtitle"`
}
type ProductList struct {
	Products []Fruit `json:"products"`
}

type Fruit struct {
	ID             string    `json:"id"`
	Gender         []string  `json:"gender"`
	Images         []string  `json:"images"`
	Reviews        []Review  `json:"reviews"`
	Publish        string    `json:"publish"`
	Ratings        []Rating  `json:"ratings"`
	Category       string    `json:"category"`
	Available      int       `json:"available"`
	PriceSale      float64   `json:"priceSale"`
	Taxes          float64   `json:"taxes"`
	Quantity       int       `json:"quantity"`
	InventoryType  string    `json:"inventoryType"`
	Tags           []string  `json:"tags"`
	Code           string    `json:"code"`
	Description    string    `json:"description"`
	SKU            string    `json:"sku"`
	CreatedAt      time.Time `json:"createdAt"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	CoverURL       string    `json:"coverUrl"`
	Colors         []string  `json:"colors"`
	TotalRatings   float64   `json:"totalRatings"`
	TotalSold      int       `json:"totalSold"`
	TotalReviews   int       `json:"totalReviews"`
	NewLabel       Label     `json:"newLabel"`
	SaleLabel      Label     `json:"saleLabel"`
	Sizes          string    `json:"sizes"`
	SubDescription string    `json:"subDescription"`
}

type Review struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	PostedAt    time.Time `json:"postedAt"`
	Comment     string    `json:"comment"`
	IsPurchased bool      `json:"isPurchased"`
	Rating      float64   `json:"rating"`
	AvatarURL   string    `json:"avatarUrl"`
	Helpful     int       `json:"helpful"`
	Attachments []string  `json:"attachments"`
}

type Rating struct {
	Name        string `json:"name"`
	StarCount   int    `json:"starCount"`
	ReviewCount int    `json:"reviewCount"`
}

type Label struct {
	Enabled bool   `json:"enabled"`
	Content string `json:"content"`
}
