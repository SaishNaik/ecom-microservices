package repo

import (
	"context"
	"github.com/SaishNaik/ecom_microservices/product/domain"
)

type ProductRepo interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
}
