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
	product.GET("", ctx.HandlerFunc(p.FindAllProducts))

}
func (p *Product) FindAllProducts(c *gin.Context) error {
	resp, err := p.ProductSvc.FindAllProduct(c)
	if err != nil {
		return ctx.Error(c, 500, err.Error())
	}
	return ctx.Success(c, resp)
}
