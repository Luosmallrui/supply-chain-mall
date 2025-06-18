package controller

import (
	"github.com/gin-gonic/gin"
	"supply-chain-mall/svc"
)

type Product struct {
	ProductSvc svc.IProductService
}

func (p *Product) RegisterRouter(r gin.IRouter) {
	product := r.Group("/")
	product.GET("/", p.FindAllProducts)

}
func (p *Product) FindAllProducts(c *gin.Context) {
	p.ProductSvc.CreateProduct()

}
