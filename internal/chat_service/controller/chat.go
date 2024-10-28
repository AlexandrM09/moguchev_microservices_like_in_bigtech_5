package controller

import (
	"context"

	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/chat_service"
)

func (c *Controller) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{},nil
}

func (c *Controller)Readyz(ctx context.Context, req *pb.ReadyzRequest) (*pb.ReadyzResponse, error) {
	return &pb.ReadyzResponse{},nil
}

//WriteFriend - Написать сообщение другу
func (c *Controller)WriteFriend(ctx context.Context,req *pb.WriteFriendRequest)(*pb.WriteFriendResponse,error){
	
	return &pb.WriteFriendResponse{},nil
}

//GetMessageChat  - Получить сообщение из чата с пользователем
func (c *Controller)GetMessageChat(ctx context.Context,req *pb.GetMessageChatRequest)(*pb.GetMessageChatResponse,error){
	return &pb.GetMessageChatResponse{},nil
}