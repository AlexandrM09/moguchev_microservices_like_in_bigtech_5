package controller

import (
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/usecases"
	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/chat_service"
)

// Deps - server deps
type Deps struct {
	ChatUsecase uc.Usecase
}

// Controller - реализация обработчика gRPC/REST запросов
type Controller struct {
	pb.UnimplementedChatServiceServer
	Deps
}

// New - returns *Controller
func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
