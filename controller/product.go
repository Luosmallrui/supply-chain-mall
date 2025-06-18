package controller

import (
	"github.com/gin-gonic/gin"
	"supply-chain-mall/pkg/ctx"
	"supply-chain-mall/svc"
	"supply-chain-mall/types"
)

type Product struct {
	ProductSvc svc.IProductService
}

func (p *Product) RegisterRouter(r gin.IRouter) {
	product := r.Group("/product")
	product.GET("", ctx.HandlerFunc(p.FindAllProducts))
	product.POST("", ctx.HandlerFunc(p.CreateProduct))

}
func (p *Product) FindAllProducts(c *gin.Context) error {
	resp, err := p.ProductSvc.FindAllProduct(c)
	if err != nil {
		return ctx.Error(c, 500, err.Error())
	}
	return ctx.Success(c, resp)
}

func (p *Product) CreateProduct(c *gin.Context) error {
	// 绑定请求体
	var product types.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		return ctx.Error(c, 400, "Invalid input: "+err.Error())
	}

	// 调用 Service 层创建产品
	if err := p.ProductSvc.CreateProduct(c, &product); err != nil {
		return ctx.Error(c, 500, "Failed to create product: "+err.Error())
	}

	// 成功返回
	return ctx.Success(c, "Product created successfully")
}
