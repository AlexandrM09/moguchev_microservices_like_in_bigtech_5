package main

import (
	"context"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/middleware/grpc"
	"log"

	controllers "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/controller"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/server"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	repo "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/repositories/inmemory"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/usecases"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usersRepo := repo.NewRepository(1000)

	// =========================
	// usecases
	// =========================

	usersUsecase := uc.NewUsecase(uc.Deps{
		UsersRepository: usersRepo,
	})

	// =========================
	// delivery
	// =========================

	// controller
	friendsController := controllers.New(controllers.Deps{UsersUsecase: usersUsecase})

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
		UsersServiceServer: friendsController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
