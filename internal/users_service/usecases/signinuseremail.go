package usecase

import (
	"context"
	"time"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/auth/jwt"
)



func (uc *usecase) SignInUserEmail(ctx context.Context, req models.SignInUserPostRequest) (*models.SignInUserPostResponse, error){
		
		//поиск юзера
		res,err:=uc.UsersRepository.SearchByNickname(ctx,models.SearchByNicknameRequest{Nickname: req.Nickname});
		if err!=nil{
			
			return nil, models.ErrSignInUserEmailFailed
		}
		

	// if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
	// 	return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	// }

	token, err :=jwt.GenerateToken(10*time.Hour,*res,"")
	if err != nil {
		return nil, models.ErrSignInUserEmailFailed
	}

 return &models.SignInUserPostResponse{Token: token}, nil
}