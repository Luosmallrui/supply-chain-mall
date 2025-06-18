package svc

var _ IProductService = (*ProductService)(nil)

type ProductService struct {
}

type IProductService interface {
	CreateProduct() error
}

func (s *ProductService) CreateProduct() error {
	return nil
}
