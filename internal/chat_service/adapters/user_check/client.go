package user_check

import (
	
	u "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/usecases"
	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/users_service"
)

type Client struct {
	grpcClient pb.UsersServiceClient
}

var (
	_ u.UsersSystem = (*Client)(nil)
)

func NewClient(client pb.UsersServiceClient) *Client {
	return &Client{
		grpcClient: client,
	}
}
