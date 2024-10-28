package inmemory

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/chat_service/models"
)

func (r *Repository)WriteFriend(ctx context.Context, req models.WriteFriendRequest) (*models.WriteFriendResponse,error){
	return &models.WriteFriendResponse{},nil
}