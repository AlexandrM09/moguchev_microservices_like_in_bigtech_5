package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)


func (uc *usecase)SearchByNickname(ctx context.Context, req models.SearchByNicknameRequest) (*models.SearchByNicknameResponse, error){
		
		//поиск юзера в бд
		res,err:=uc.UsersRepository.SearchByNickname(ctx,req);
		if err!=nil{
			
			return nil, models.ErrSearchByNicknameFailed
		}
		//todo dto
 return res, nil
}