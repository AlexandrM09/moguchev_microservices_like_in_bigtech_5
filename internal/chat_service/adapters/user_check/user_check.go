package user_check

import (
	 "context"
pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/users_service"
)

func (c *Client) Check(ctx context.Context, nickname string) error {

	 _,err:=c.grpcClient.SearchByNickname(ctx,&pb.SearchByNicknameRequest{Nickname:nickname })

	return err
}