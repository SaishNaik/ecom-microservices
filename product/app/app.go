package app

import (
	"github.com/SaishNaik/ecom_microservices/product/config"
	gen "github.com/SaishNaik/ecom_microservices/product/proto/gen/go/proto"
)

type App struct {
	Config            *config.Config
	ProductGRPCServer gen.ProductServiceServer
}

func NewApp(cfg *config.Config, productGrpcServer gen.ProductServiceServer) *App {
	return &App{
		Config:            cfg,
		ProductGRPCServer: productGrpcServer,
	}
}
