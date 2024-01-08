package controller

import (
	"context"
	"fmt"
	gen "github.com/SaishNaik/ecom_microservices/product/proto/gen/go/proto"
	"github.com/SaishNaik/ecom_microservices/product/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

type productGrpcServer struct {
	gen.UnimplementedProductServiceServer
	productService service.ProductService
}

var _ gen.ProductServiceServer = (*productGrpcServer)(nil)

var ProductGrpcServerSet = wire.NewSet(NewProductGrpcServer)

func NewProductGrpcServer(grpcServer *grpc.Server, productService service.ProductService) gen.ProductServiceServer {
	grpcSvc := productGrpcServer{productService: productService}
	gen.RegisterProductServiceServer(grpcServer, &grpcSvc)
	return &grpcSvc
}

func (p *productGrpcServer) GetAllProducts(ctx context.Context, request *gen.GetAllProductsRequest) (*gen.GetAllProductsResponse, error) {
	res := &gen.GetAllProductsResponse{}

	results, err := p.productService.GetAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error From GRPCServer GetAllProducts %w", err)
	}

	for _, product := range results {
		res.Products = append(res.Products, &gen.Product{
			Name:  product.Name,
			Price: product.Price,
			Image: product.Image,
			Sku:   product.Sku,
		})
	}

	return res, nil
}

//package router
//
//import (
//	"context"
//
//	"github.com/google/wire"
//	"github.com/pkg/errors"
//	"github.com/thangchung/go-coffeeshop/internal/product/usecases/products"
//	"github.com/thangchung/go-coffeeshop/proto/gen"
//	"golang.org/x/exp/slog"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//)
//
//var _ gen.ProductServiceServer = (*productGRPCServer)(nil)
//
//var ProductGRPCServerSet = wire.NewSet(NewProductGRPCServer)
//
//type productGRPCServer struct {
//	gen.UnimplementedProductServiceServer
//	uc products.UseCase
//}
//
//func NewProductGRPCServer(
//	grpcServer *grpc.Server,
//	uc products.UseCase,
//) gen.ProductServiceServer {
//	svc := productGRPCServer{
//		uc: uc,
//	}
//
//	gen.RegisterProductServiceServer(grpcServer, &svc)
//
//	reflection.Register(grpcServer)
//
//	return &svc
//}
//
//func (g *productGRPCServer) GetItemTypes(
//	ctx context.Context,
//	request *gen.GetItemTypesRequest,
//) (*gen.GetItemTypesResponse, error) {
//	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")
//
//	res := gen.GetItemTypesResponse{}
//
//	results, err := g.uc.GetItemTypes(ctx)
//	if err != nil {
//		return nil, errors.Wrap(err, "productGRPCServer-GetItemTypes")
//	}
//
//	for _, item := range results {
//		res.ItemTypes = append(res.ItemTypes, &gen.ItemTypeDto{
//			Name:  item.Name,
//			Type:  int32(item.Type),
//			Price: item.Price,
//			Image: item.Image,
//		})
//	}
//
//	return &res, nil
//}
//
//func (g *productGRPCServer) GetItemsByType(
//	ctx context.Context,
//	request *gen.GetItemsByTypeRequest,
//) (*gen.GetItemsByTypeResponse, error) {
//	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemsByType", "item_types", request.ItemTypes)
//
//	res := gen.GetItemsByTypeResponse{}
//
//	results, err := g.uc.GetItemsByType(ctx, request.ItemTypes)
//	if err != nil {
//		return nil, errors.Wrap(err, "productGRPCServer-GetItemsByType")
//	}
//
//	for _, item := range results {
//		res.Items = append(res.Items, &gen.ItemDto{
//			Type:  int32(item.Type),
//			Price: item.Price,
//		})
//	}
//
//	return &res, nil
//}
