package models

import "time"

//WriteFriendRequest - запрос написать сообщение другу
type  WriteFriendRequest struct {
	//никнейм - уникальный
	Nickname string
	//никнейм друга - уникальный
	NicknameFriend string
	//сообщение
	Message string
  }
  
  //WriteFriendResponse - ответ
  type WriteFriendResponse struct {
	//номер сообщения
	Position uint64
  }
  
  //GetMessageChatRequest  - Получить сообщение из чата с пользователем
 type GetMessageChatRequest struct{
	//никнейм - уникальный
	Nickname string
	//никнейм друга - уникальный
	NicknameFriend string
	//последняя прочитанная позиция
	LastPosition uint64 
  }
  
  //Сообщение в чате
 type OneMessages struct{
	//никнейм - уникальный
	Nickname string 
	//сообщение в чате
	Message string 
	//номер последнего сообщения
	Position uint64 
	//Время сообщения
	CreatedAt time.Time
  }
  
  //GetMessageChatResponse -ответ cписок сообщений
  type GetMessageChatResponse struct{
	//Список сообщений
	Messages []OneMessages
  }