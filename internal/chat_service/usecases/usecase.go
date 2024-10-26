package usecase

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
	//orders_models "lesson_08/orders_management_system/internal/app/usecases/orders/models"
	//"lesson_08/orders_management_system/pkg/pagination"
)

// Usecase - port (первичный)
//
// ...
type Usecase interface {
	//WriteFriend - Написать сообщение другу
	WriteFriend(ctx context.Context, req models.WriteFriendRequest) (*models.WriteFriendResponse, error)
	//GetMessageChat  - Получить сообщение из чата с пользователем
	GetMessageChat(ctx context.Context, req models.GetMessageChatRequest) (*models.GetMessageChatResponse, error)
}

type (
	// Users - port (вторичный)
	UsersSystem interface {
		Check(ctx context.Context, Nickname string) error
	}

	// ChatRepository - port (вторичный)
	ChatRepository interface {
		WriteFriend(ctx context.Context, req models.WriteFriendRequest) (*models.WriteFriendResponse,error)
		GetMessageChat(ctx context.Context, req models.GetMessageChatRequest) (*models.GetMessageChatResponse,error)
	}
)

// Deps -
type Deps struct {
	// Adapters
	Users              UsersSystem
	ChatRepository ChatRepository
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	Deps
}

// бизнес сервис
func NewUsecase(d Deps) *usecase {
	return &usecase{Deps: d}
}
