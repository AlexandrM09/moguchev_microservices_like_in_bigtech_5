package inmemory

import (
	"sync"


	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/usecases"
	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/repositories/dto"
)
type key struct{
	//никнейм - уникальный
	Nickname string
	//никнейм друга - уникальный
	// NicknameFriend string
}

  

type Repository struct {
	storage map[key] dto.UsersRepo
	mx      sync.Mutex
}

var (
	_ uc.UsersRepository = (*Repository)(nil)
)

func NewRepository(cap int) *Repository {
	return &Repository{storage: make(map[key]dto.UsersRepo, cap)}
}

func toUsersRepo(u models.SearchByNicknameResponse)dto.UsersRepo{
	return dto.UsersRepo{
		//Id пользователя
		UserId:u.UserId,
		//никнейм - уникальный
		Nickname:u.Nickname,
		//информация о себе
		Description:u.Description,
		//аватарка
		Avatar:u.Avatar,
		//Email
		Email:u.Email,
		//Provider
		Provider:u.Provider,
		//Время создания
		CreatedAt:u.CreatedAt,
		//Время последнего обновления
		UpdatedAt:u.UpdatedAt,
	}
}