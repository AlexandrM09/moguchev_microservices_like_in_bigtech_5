package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)


func (uc *usecase)Confirm(ctx context.Context, req models.ConfirmRequest) (*models.ConfirmResponse, error){
	//Проверка существования ника запрашивающего
	if  err:=uc.Users.Check(ctx,req.Nickname);err!=nil{
		return nil, models.ErrNotFoundNickname
	}
	//Проверка существования друга
	if  err:=uc.Users.Check(ctx,req.NicknameFriend);err!=nil{
		return nil, models.ErrNotFoundFriend
	}
	//Запись сообщения бд
	res,err:=uc.FriendRepository.Confirm(ctx,req);
	if err!=nil{
		return nil, models.ErrRemoveFailed
	}
	return res,nil
}