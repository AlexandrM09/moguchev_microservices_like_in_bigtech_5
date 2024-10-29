package controller

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	//uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/usecases"
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
		//todo pb->model
		_,err:=c.UsersUsecase.SignUpUserEmail(ctx,models.SignUpUserEmailRequest{})
		//todo model->pb	
	return &pb.SignUpUserEmailResponse{},err
}

//SignInUser - вход/авторизация (по почте и паролю)
func (c *Controller) SignInUserEmail(ctx context.Context, req *pb.SignInUserPostRequest) (*pb.SignInUserPostResponse, error){
	//todo pb->model
	_,err:=c.UsersUsecase.SignInUserEmail(ctx,models.SignInUserPostRequest{})
	//todo model->pb	
	return &pb.SignInUserPostResponse{},err
}

//SignInUserOauth - вход/авторизация (Oauth), регистрация если нет такого пользователя
func (c *Controller) SignInUserOauth(ctx context.Context, req *pb.SignInUserOauthRequest) (*pb.SignInUserOauthResponse, error){
	//todo pb->model
	_,err:=c.UsersUsecase.SignInUserOauth(ctx,models.SignInUserOauthRequest{})
	//todo model->pb
	return &pb.SignInUserOauthResponse{},err
}

//UpdateProfile - редактирование профиля пользователя (никнейм - уникальный, информация о себе, аватарка)
func (c *Controller) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error){
	//todo pb->model
	_,err:=c.UsersUsecase.UpdateProfile(ctx,models.UpdateProfileRequest{})
	//todo model->p
	return &pb.UpdateProfileResponse{},err
}

//SearchByNickname- поиск пользователей по никнейму
func (c *Controller) SearchByNickname(ctx context.Context, req *pb.SearchByNicknameRequest) (*pb.SearchByNicknameResponse, error){
	//todo pb->model
	_,err:=c.UsersUsecase.SearchByNickname(ctx,models.SearchByNicknameRequest{})
	//todo model->p
	return &pb.SearchByNicknameResponse{},err
}