package inmemory

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/friends_service/models"
)

func (r *Repository)GetList(ctx context.Context, req models.GetListRequest) (*models.GetListResponse, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	//проверка на существование
	keyA:=key{Nickname:req.Nickname}
	v,ok:=r.storage[keyA]
	
	if !ok {
		return nil,models.ErrGetListFailed
	}
	
	res:=models.GetListResponse{}
	for _,e:=range v{
      if e.Status{
		res.Confirmed=append(res.Confirmed, e.NicknameFriend)
		continue
	  }    
	  res.NotConfirmed=append(res.NotConfirmed, e.NicknameFriend)
	}
	return &res,nil
}