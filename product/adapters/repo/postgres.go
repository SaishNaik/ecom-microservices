package repo

import (
	"context"
	"fmt"
	"github.com/SaishNaik/ecom_microservices/product/domain"
	"github.com/SaishNaik/ecom_microservices/product/repo"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type postgresRepo struct {
	pool *pgxpool.Pool
}

var _ repo.ProductRepo = (*postgresRepo)(nil)

var ProductRepoSet = wire.NewSet(NewPostgresRepo)

func NewPostgresRepo(dns string) (repo.ProductRepo, error) {
	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		slog.Error("Postgres failed with error", err)
		return nil, fmt.Errorf("error from NewPostgresRepo while creating pool: %w", err)
	}
	return &postgresRepo{
		pool: pool,
	}, nil
}

func (p *postgresRepo) GetAllProducts(ctx context.Context) ([]domain.Product, error) {

	var products []domain.Product
	query := "SELECT * from products"
	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error from GetAllProducts while executing query: %w", err)
	}
	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product)
		if err != nil {
			return nil, fmt.Errorf("error from GetAllProducts while scanning: %w", err)
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error from GetAllProducts while reading: %w", rows.Err())
	}

	return products, err

	//return []domain.Product{{
	//	Sku:   "1222",
	//	Name:  "Product1",
	//	Price: 50,
	//	Image: "http://1.com",
	//},
	//}, nil
}
