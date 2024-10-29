package inmemory

import (
	"context"
	"fmt"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)

func (r *Repository)UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) (*models.UpdateProfileResponse, error){
	r.mx.Lock()
	defer r.mx.Unlock()
	key:=key{Nickname:req.Nickname}
	//read
	v,ok:=r.storage[key]
	if !ok{
		return nil,fmt.Errorf("user not exist")
	}
	//change
	v.Nickname=req.Nickname
	v.Description=req.Description
	v.Avatar=req.Avatar
	//save
	r.storage[key]=v
	return &models.UpdateProfileResponse{},nil
}