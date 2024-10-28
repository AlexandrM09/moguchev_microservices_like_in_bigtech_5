package usecase

import (
	"context"
	//"errors"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)

func (uc *usecase) SearchByNickname(ctx context.Context, req models.SearchByNicknameRequest) (*models.SearchByNicknameResponse, error) {

	//поиск юзера в бд
	res, err := uc.UsersRepository.GetUser(ctx, req)
	if err != nil {

		return nil, models.ErrSearchByNicknameFailed
	}
	//todo dto
	return &models.SearchByNicknameResponse{
		//Id пользователя
		UserId: res.UserId,
		//никнейм - уникальный
		Nickname: res.UserId,
		//информация о себе
		Description: res.UserId,
		//аватарка
		Avatar: res.UserId,
		//Email
		Email: res.UserId,
		//Provider
		Provider: res.UserId,
		//Время создания
		CreatedAt: res.CreatedAt,
		//Время последнего обновления
		UpdatedAt: res.UpdatedAt,
	}, nil
}
