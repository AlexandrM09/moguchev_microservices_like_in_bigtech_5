package controller

import (
	"context"

	pb "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/api/users_service"
)

func (c *Controller) Healthz(ctx context.Context, req *pb.HealthzRequest) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{}, nil
}

func (c *Controller) Readyz(ctx context.Context, req *pb.ReadyzRequest) (*pb.ReadyzResponse, error) {
	return &pb.ReadyzResponse{}, nil
}

//SignUpUser - регистрация пользователя (по почте и паролю)
func (c *Controller) SignUpUserEmail(ctx context.Context, req *pb.SignUpUserEmailRequest) (*pb.SignUpUserEmailResponse, error){
	return &pb.SignUpUserEmailResponse{},nil
}

//SignInUser - вход/авторизация (по почте и паролю)
func (c *Controller) SignInUserEmail(ctx context.Context, req *pb.SignInUserPostRequest) (*pb.SignInUserPostResponse, error){
	return &pb.SignInUserPostResponse{},nil
}

//SignInUserOauth - вход/авторизация (Oauth), регистрация если нет такого пользователя
func (c *Controller) SignInUserOauth(ctx context.Context, req *pb.SignInUserOauthRequest) (*pb.SignInUserOauthResponse, error){
	return &pb.SignInUserOauthResponse{},nil
}

//UpdateProfile - редактирование профиля пользователя (никнейм - уникальный, информация о себе, аватарка)
func (c *Controller) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error){
	return &pb.UpdateProfileResponse{},nil
}

//SearchByNickname- поиск пользователей по никнейму
func (c *Controller) SearchByNickname(ctx context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error){
	return &pb.SearchByNicknameResponse{},nil
}