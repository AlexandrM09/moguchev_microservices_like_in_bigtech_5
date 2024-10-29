package inmemory

import (

	"sync"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/usecases"
)
type key struct{
	//никнейм - уникальный
	Nickname string
	//никнейм друга - уникальный
	NicknameFriend string
}

type Repository struct {
	storage map[key] chat
	mx      sync.Mutex
}

var (
	_ uc.ChatRepository = (*Repository)(nil)
)

func NewRepository(cap int) *Repository {
	return &Repository{storage: make(map[key]chat, cap)}
}

type chat struct {
	*models.GetMessageChatResponse
}


