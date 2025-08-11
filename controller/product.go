package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"supply-chain-mall/pkg/ctx"
	"supply-chain-mall/svc"
	"supply-chain-mall/types"
	"time"
)

type Product struct {
	ProductSvc svc.IProductService
}

func (p *Product) RegisterRouter(r gin.IRouter) {
	product := r.Group("/product")
	product.GET("", ctx.HandlerFunc(p.FindAllProducts))
	product.POST("", ctx.HandlerFunc(p.CreateProduct))
	g := r.Group("/fruit")
	g.GET("/list", GetFruits)
	g.GET("/details", GetFruitsDetail)
}

func GetFruits(c *gin.Context) {
	filePath := "a.json"

	// 读取文件
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 定义变量接收解析结果
	var productList types.ProductList

	// 解析 JSON
	err = json.Unmarshal(jsonBytes, &productList)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	// 使用解析后的数据
	fmt.Printf("共解析到 %d 个产品\n", len(productList.Products))
	for _, p := range productList.Products {
		fmt.Printf("产品 ID: %s 名称: %s 创建时间: %s\n", p.ID, p.Name, p.CreatedAt.Format(time.RFC3339))
	}
	c.JSON(200, gin.H{"data": productList})
}

func GetFruitsDetail(c *gin.Context) {
	filePath := "b.json"

	// 读取文件
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 定义变量接收解析结果
	var productList types.Fruit

	// 解析 JSON
	err = json.Unmarshal(jsonBytes, &productList)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}
	c.JSON(200, gin.H{"data": productList})
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
