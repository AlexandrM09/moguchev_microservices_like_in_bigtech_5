package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)


func (uc *usecase) UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) (*models.UpdateProfileResponse, error){
		
		//обновление профиля юзера в бд
		res,err:=uc.UsersRepository.UpdateProfile(ctx,req);
		if err!=nil{
			
			return nil, models.ErrUpdateProfileFailed
		}
 return res, nil
}