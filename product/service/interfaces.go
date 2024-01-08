package service

import (
	"context"
	"github.com/SaishNaik/ecom_microservices/product/domain"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
}
