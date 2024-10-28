package inmemory

import (
	"context"
	"slices"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)

func (r *Repository)Confirm(ctx context.Context, req models.ConfirmRequest) (*models.ConfirmResponse, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	//проверка на существование
	keyA:=key{Nickname:req.Nickname}
	v,ok:=r.storage[keyA]
	f1:=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: false}
	if !ok||(!slices.Contains(v,f1)) {
		return nil,models.ErrConfirmFailed
	}
	v[slices.Index(v,f1)]=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: true}
    r.storage[keyA]=v
	
	return &models.ConfirmResponse{},nil
}