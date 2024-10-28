package inmemory

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
)

func (r *Repository)GetMessageChat(ctx context.Context, req models.GetMessageChatRequest) (*models.GetMessageChatResponse,error){
	return &models.GetMessageChatResponse{},nil
}