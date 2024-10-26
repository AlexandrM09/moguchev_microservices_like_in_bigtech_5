package models


//AddFriendRequest - запрос на добавить пользователя в друзья
type AddRequest struct{
	//никнейм - уникальный
	Nickname string
  }
  
  //AddFriendResponse - ответ
type AddResponse struct{}
  
  //RemoveFriendRequest -запрос убрать пользователя из друзей
  type RemoveRequest struct{
	//никнейм - уникальный
	Nickname string
  }
  
  //RemoveFriendResponse - ответ
  type RemoveResponse struct{}
  
  //ConfirmRequest -запрос подтвердить или отклонить запрос на дружбу
  type ConfirmRequest struct{
	//true - подтвердить, false - отклонить
	Confirm bool
  }
  
  //ConfirmResponse - ответ
  type ConfirmResponse struct{}
  
  //GetListRequest - запрос на просмотр списка своих друзей (подтвердивших и не подтвердивших еще)
  type GetListRequest struct{}
  
  //GetListResponse ответ просмотр списка своих друзей (подтвердивших и не подтвердивших еще)
  type GetListResponse struct{
	//список подтвердивших
	Confirmed []string
	//список не подтвердивших
	NotConfirmed []string
  }