package usecase

import (
	"context"
	"time"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/pkg/auth/jwt"
)



func (uc *usecase) SignInUserOauth(ctx context.Context, req models.SignInUserOauthRequest) (*models.SignInUserOauthResponse, error){
		
	//todo доделать

	token, err :=jwt.GenerateToken(10*time.Hour,struct{}{},"")
	if err != nil {
		return nil, models.ErrSignInUserOauthFailed
	}

 return &models.SignInUserOauthResponse{Token: token}, nil
}