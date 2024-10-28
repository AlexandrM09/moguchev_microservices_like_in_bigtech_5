package inmemory

import (
	"context"
	"slices"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)

func (r *Repository)Add(ctx context.Context, req models.AddRequest) (*models.AddResponse, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	//проверка на существование
	keyA:=key{Nickname:req.Nickname}
	v,ok:=r.storage[keyA]
	f1:=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: false}
	f2:=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: true}
	if ok&&v!=nil&&(slices.Contains(v,f1)||slices.Contains(v,f2)) {
		return nil,models.ErrAddFailed
	}
	v=append(v, f1)
    r.storage[keyA]=v
	return &models.AddResponse{},nil
}