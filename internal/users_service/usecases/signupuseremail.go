package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)



func (uc *usecase) SignUpUserEmail(ctx context.Context, req models.SignUpUserEmailRequest) (*models.SignUpUserEmailResponse, error){
		
		//создание юзера в бд
		res,err:=uc.UsersRepository.SignUpUserEmail(ctx,req);
		if err!=nil{
			//todo дополнительно проверить на ошибку ErrSignUpUserFailedUserExist 
			return nil, models.ErrSignUpUserFailed
		}
 return res, nil
}