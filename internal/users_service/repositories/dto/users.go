package dto

import "time"

//usersRepo - структура для хранения инфы о юзере
type UsersRepo struct{
	//Id пользователя
	UserId string
	//никнейм - уникальный
	Nickname string
	//информация о себе
	Description string
	//аватарка
	Avatar string
	//Email
	Email string
	//Provider
	Provider string
	//Время создания
	CreatedAt time.Time
	//Время последнего обновления
	UpdatedAt time.Time
	//HashPassword
	HashPassword string
}
//
type CreateUserResponse struct{
	//Id пользователя
	UserId string
	//никнейм - уникальный
	Nickname string
}


