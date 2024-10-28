package inmemory

import (
	"context"
	"fmt"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/repositories/dto"
)

func (r *Repository)GetUser(ctx context.Context, req models.SearchByNicknameRequest) (*dto.UsersRepo, error){
	r.mx.Lock()
	defer r.mx.Unlock()
	key:=key{Nickname:req.Nickname}
	//read
	v,ok:=r.storage[key]
	if !ok{
		return nil,fmt.Errorf("user not exist")
	}
	//dto
	u:=dto.UsersRepo{
		//Id пользователя
		UserId:v.UserId,
		//никнейм - уникальный
		Nickname:v.Nickname,
		//информация о себе
		Description:v.Description,
		//аватарка
		Avatar:v.Avatar,
		//Email
		Email:v.Email,
		//Provider
		Provider:v.Provider,
		//Время создания
		CreatedAt:v.CreatedAt,
		//Время последнего обновления
		UpdatedAt:v.UpdatedAt,
		HashPassword: v.HashPassword,
	}
	
	return &u,nil
}