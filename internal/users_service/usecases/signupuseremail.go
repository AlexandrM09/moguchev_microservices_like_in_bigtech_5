package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)



func (uc *usecase) SignUpUserEmail(ctx context.Context, req models.SignUpUserEmailRequest) (*models.SignUpUserEmailResponse, error){
		//создаем hash
        hash:="" 
		//создание юзера в бд
		_,err:=uc.UsersRepository.CreateUser(ctx,models.SignUpUserEmailRequest{},hash,"oauth");
		if err!=nil{
			//todo дополнительно проверить на ошибку ErrSignUpUserFailedUserExist 
			return nil, models.ErrSignUpUserFailed
		}
 return &models.SignUpUserEmailResponse{}, nil
}