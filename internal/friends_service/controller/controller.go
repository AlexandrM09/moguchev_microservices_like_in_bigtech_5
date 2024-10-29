package controller

import (
	//"lesson_06/orders_management_system/internal/app/usecases/orders"
	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/friends_service"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/usecases"
)

// Deps - server deps
type Deps struct {
	FriendUsecase uc.Usecase
}

// Controller - реализация обработчика gRPC/REST запросов
type Controller struct {
	pb.UnimplementedFriendsServiceServer
	Deps
}

// New - returns *Controller
func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
