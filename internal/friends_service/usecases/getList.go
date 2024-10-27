package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)


func (uc *usecase)GetList(ctx context.Context, req models.GetListRequest) (*models.GetListResponse, error){
	//Проверка существования ника запрашивающего
	if  err:=uc.Users.Check(ctx,req.Nickname);err!=nil{
		return nil, models.ErrNotFoundNickname
	}
	//Запись сообщения бд
	res,err:=uc.FriendRepository.GetList(ctx,req);
	if err!=nil{
		return nil, models.ErrGetListFailed
	}
	return res,nil
}