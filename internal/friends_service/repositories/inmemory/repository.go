package inmemory

import (

	"sync"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
	uc "github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/usecases"
)
type key struct{
	//никнейм - уникальный
	Nickname string
	//никнейм друга - уникальный
	// NicknameFriend string
}

  //ConfirmRequest -запрос подтвердить или отклонить запрос на дружбу
  type frieandsRepo struct{
    
 	//никнейм друга - уникальный
  	NicknameFriend string
	//true - подтвердить, false - отклонить
	Status bool
  }

type Repository struct {
	storage map[key] []frieandsRepo
	mx      sync.Mutex
}

var (
	_ uc.FriendRepository = (*Repository)(nil)
)

func NewRepository(cap int) *Repository {
	return &Repository{storage: make(map[key][]frieandsRepo, cap)}
}

type chat struct {
	*models.GetMessageChatResponse
}