//go:build wireinject
// +build wireinject

package app

import (
	"github.com/SaishNaik/ecom_microservices/product/adapters/repo"
	"github.com/SaishNaik/ecom_microservices/product/config"
	"github.com/SaishNaik/ecom_microservices/product/controller"
	"github.com/SaishNaik/ecom_microservices/product/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
	dns string,
) (*App, error) {

	panic(wire.Build(
		NewApp,
		controller.ProductGrpcServerSet,
		service.ProductServiceSet,
		repo.ProductRepoSet,
	))
}
