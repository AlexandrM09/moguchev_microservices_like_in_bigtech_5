package main

import (
	"context"
	"log"

	adapter "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/adapters/user_check"
	controllers "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/controller"
	repo "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/repositories/inmemory"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/server"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/usecases"
	usersClient "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/users_service"
	middleware_grpc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/middleware/grpc"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	//todo вынести в переменные окружения
	conn, err := grpc.NewClient("localhost:8082",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	//todo add client interceptor to auth
	usersAdapter := adapter.NewClient(
		usersClient.NewUsersServiceClient(conn),
	)
	// =========================
	// repository

	chatRepo := repo.NewRepository(1000)

	// =========================
	// usecases
	// =========================

	chatUsecase := uc.NewUsecase(uc.Deps{
		Users:          usersAdapter,
		ChatRepository: chatRepo,
	})

	// =========================
	// delivery
	// =========================

	// controller
	chatController := controllers.New(controllers.Deps{ChatUsecase: chatUsecase})

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
		ChatServiceServer: chatController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
	//todo gracefull shotdown
}
