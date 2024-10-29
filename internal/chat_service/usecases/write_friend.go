package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
)


func (uc *usecase)WriteFriend(ctx context.Context, req models.WriteFriendRequest) (*models.WriteFriendResponse, error){
	//Проверка существования ника запрашивающего
	if  err:=uc.Users.Check(ctx,req.Nickname);err!=nil{
		return nil, models.ErrNotFoundNickname
	}
	//Проверка существования друга
	if  err:=uc.Users.Check(ctx,req.NicknameFriend);err!=nil{
		return nil, models.ErrNotFoundFriend
	}
	//todo добавить проверку что req.Nickname и req.NicknameFriend друзья 
	//Запись сообщения бд
	res,err:=uc.ChatRepository.WriteFriend(ctx,req);
	if err!=nil{
		return nil, models.ErrWriteFriendFailed
	}
	return res,nil
}