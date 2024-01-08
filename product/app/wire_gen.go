// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/SaishNaik/ecom_microservices/product/adapters/repo"
	"github.com/SaishNaik/ecom_microservices/product/config"
	"github.com/SaishNaik/ecom_microservices/product/controller"
	"github.com/SaishNaik/ecom_microservices/product/service"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, grpcServer *grpc.Server, dns string) (*App, error) {
	productRepo, err := repo.NewPostgresRepo(dns)
	if err != nil {
		return nil, err
	}
	productService := service.NewService(productRepo)
	productServiceServer := controller.NewProductGrpcServer(grpcServer, productService)
	app := NewApp(cfg, productServiceServer)
	return app, nil
}
