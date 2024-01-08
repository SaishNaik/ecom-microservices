package service

import (
	"context"
	"fmt"
	"github.com/SaishNaik/ecom_microservices/product/domain"
	"github.com/SaishNaik/ecom_microservices/product/repo"
	"github.com/google/wire"
)

var _ ProductService = (*productService)(nil)

var ProductServiceSet = wire.NewSet(NewService)

type productService struct {
	repo repo.ProductRepo
}

func NewService(productRepo repo.ProductRepo) ProductService {
	return &productService{
		repo: productRepo,
	}
}

func (p *productService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := p.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("Service::GetAllProducts %w", err)
	}
	return products, err
}
