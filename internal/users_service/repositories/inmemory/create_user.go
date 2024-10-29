package inmemory

import (
	"context"
	"fmt"
	"time"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/repositories/dto"
	"github.com/google/uuid"
)

func (r *Repository)CreateUser(ctx context.Context, req  models.SignUpUserEmailRequest,hash string,provider string) (*dto.CreateUserResponse, error){
	r.mx.Lock()
	defer r.mx.Unlock()
	key:=key{Nickname:req.Nickname}
	//
	if _,ok:=r.storage[key];ok{
		return nil,fmt.Errorf("user exist")
	}
	//dto
	u:=dto.UsersRepo{
		//Id пользователя
		UserId:uuid.New().String(),
		//никнейм - уникальный
		Nickname:req.Nickname,
		//информация о себе
		Description:req.Description,
		//аватарка
		Avatar:req.Avatar,
		//Email
		Email:req.Email,
		//Provider
		Provider:provider,
		//Время создания
		CreatedAt:time.Now(),
		//Время последнего обновления
		UpdatedAt:time.Now(),
		HashPassword: hash,
	}
	//save
	r.storage[key]=u
	return &dto.CreateUserResponse{UserId:u.UserId,Nickname: u.Nickname},nil
}