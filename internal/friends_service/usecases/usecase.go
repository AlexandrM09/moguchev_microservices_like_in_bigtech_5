package usecase

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)

// Usecase - port (первичный)
//
// ...
type Usecase interface {
	//Add - Добавить пользователя в друзья
	Add(ctx context.Context, req models.AddRequest) (*models.AddResponse, error)
	//Remove - убрать пользователя из друзей
	Remove(ctx context.Context, req models.RemoveRequest) (*models.RemoveResponse, error)
	//Confirm - Подтвердить или отклонить запрос на дружбу
	Confirm(ctx context.Context, req models.ConfirmRequest) (*models.ConfirmResponse, error)
	//GetList - Просмотр списка своих друзей (подтвердивших и не подтвердивших еще)
	GetList(ctx context.Context, req models.GetListRequest) (*models.GetListResponse, error)
}

type (
	// Users - port (вторичный)
	UsersSystem interface {
		Check(ctx context.Context, Nickname string) error
	}

	// FriendRepository - port (вторичный)
	FriendRepository interface {
		//Add - Добавить пользователя в друзья
		Add(ctx context.Context, req models.AddRequest) (*models.AddResponse, error)
		//Remove - убрать пользователя из друзей
		Remove(ctx context.Context, req models.RemoveRequest) (*models.RemoveResponse, error)
		//Confirm - Подтвердить или отклонить запрос на дружбу
		Confirm(ctx context.Context, req models.ConfirmRequest) (*models.ConfirmResponse, error)
		//GetList - Просмотр списка своих друзей (подтвердивших и не подтвердивших еще)
		GetList(ctx context.Context, req models.GetListRequest) (*models.GetListResponse, error)
	}
)

// Deps -
type Deps struct {
	// Adapters
	Users            UsersSystem
	FriendRepository FriendRepository
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	Deps
}

// бизнес сервис
func NewUsecase(d Deps) *usecase {
	return &usecase{Deps: d}
}
