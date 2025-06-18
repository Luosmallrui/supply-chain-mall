package controller

import (
	"github.com/gin-gonic/gin"
	"supply-chain-mall/pkg/ctx"
	"supply-chain-mall/svc"
)

type Product struct {
	ProductSvc svc.IProductService
}

func (p *Product) RegisterRouter(r gin.IRouter) {
	product := r.Group("/product")
	product.GET("/", p.FindAllProducts)

}
func (p *Product) FindAllProducts(c *gin.Context) {
	p.ProductSvc.CreateProduct()

	ctx.Error(c, 400, "failed")
}
