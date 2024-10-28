package inmemory

import (
	"context"
	"slices"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)

func (r *Repository)Remove(ctx context.Context, req models.RemoveRequest) (*models.RemoveResponse, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	//проверка на существование
	keyA:=key{Nickname:req.Nickname}
	v,ok:=r.storage[keyA]
	f1:=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: false}
	f2:=frieandsRepo{NicknameFriend: req.NicknameFriend, Status: true}
	if !ok&&(slices.Contains(v,f1)||slices.Contains(v,f2)) {
		return nil,models.ErrRemoveFailed
	}
	ind:=slices.Index(v,f1)
	if ind==-1{
		ind=slices.Index(v,f2)
	}
	v=slices.Delete(v,ind,ind)
    r.storage[keyA]=v
	return &models.RemoveResponse{},nil
}