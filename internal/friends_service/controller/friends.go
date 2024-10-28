package controller

import (
	"context"

	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/friends_service"
)

func (c *Controller) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{}, nil
}

func (c *Controller) Readyz(ctx context.Context, req *pb.ReadyzRequest) (*pb.ReadyzResponse, error) {
	return &pb.ReadyzResponse{}, nil
}

//Add - Добавить пользователя в друзья
func (c *Controller)Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error){
	return &pb.AddResponse{},nil
}
//Remove - убрать пользователя из друзей
func (c *Controller)Remove(ctx context.Context, req *pb.RemoveRequest) (*pb.RemoveResponse, error){
	return &pb.RemoveResponse{},nil
}
//Confirm - Подтвердить или отклонить запрос на дружбу
func (c *Controller)Confirm(ctx context.Context, req *pb.ConfirmRequest) (*pb.ConfirmResponse, error){
	return &pb.ConfirmResponse{},nil
}
//GetList - Просмотр списка своих друзей (подтвердивших и не подтвердивших еще)
func (c *Controller)GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error){
	return &pb.GetListResponse{},nil
}