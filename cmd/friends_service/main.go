package main

import (
	"context"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/middleware/grpc"
	"log"

	controllers "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/controller"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/server"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository

	//ordersRepo := orders_repository.NewRepository(1000)

	// services

	//wmsAdapter := warehouse_management_system.NewClient()

	// =========================
	// usecases
	// =========================

	//ordersUsecase := orders.NewUsecase(orders.Deps{
	//	WMS:              wmsAdapter,
	//	OrdersRepository: ordersRepo,
	//})

	// =========================
	// delivery
	// =========================

	// controller
	friendsController := controllers.New(controllers.Deps{})

	// middlewares

	validator, err := protovalidate.New(protovalidate.WithDisableLazy(false))
	if err != nil {
		log.Fatalf("server: failed to initialize validator: %s", err)
	}
	mws := []grpc.UnaryServerInterceptor{
		middleware_grpc.ErrorsUnaryServerInterceptor(),
		middleware_grpc.ValidateUnaryServerInterceptor(validator),
		middleware_grpc.AuthorizationUnaryServerInterceptor(),
	}

	// infrastructure server
	config := server.Config{
		GRPCPort:               ":8082",
		GRPCGatewayPort:        ":8080",
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.New(ctx, config, server.Contollers{
		FriendsServiceServer: friendsController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
