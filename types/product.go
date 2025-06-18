package types

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`         // 改为浮点数
	OriginalPrice int     `json:"originalPrice"` // 改为整数
	Image         string  `json:"image"`
	Brand         string  `json:"brand"`
	BrandSubtitle string  `json:"brandSubtitle"`
}
