package models

import "time"

// SignUpUserEmailRequest - запрос регистрации с почтой и паролем SignUpUserEmail(
type SignUpUserEmailRequest struct {
	//никнейм - уникальный
	Nickname string
	//информация о себе
	Description string
	//аватарка
	Avatar string
	//Email
	Email string
	//Password
	Password string
	//Role  string
	//Verified bool
	//Provider string
}

//Ответ регистрации
type SignUpUserEmailResponse struct{}

//SignInUserPostRequest - вход/авторизация (по почте и паролю)
type SignInUserPostRequest struct {
	//никнейм - уникальный
	Nickname string
	//Email
	Email string
	//Password
	Password string
}

//SignInUserPostResponse
type SignInUserPostResponse struct {
	Token string
}

//SignInUserOauthRequest - вход/авторизация (Oauth)
type SignInUserOauthRequest struct {
	//
	Code string
	//
	State string
}

//SignInUserOauthResponse
type SignInUserOauthResponse struct {
	Token string
}

//UpdateProfile - редактирование профиля пользователя (никнейм - уникальный, информация о себе, аватарка)
type UpdateProfileRequest struct {
	//никнейм - уникальный
	Nickname string
	//информация о себе
	Description string
	//аватарка
	Avatar string
}

type UpdateProfileResponse struct{}

//запрос в SearchByNickname- поиск пользователей по никнейму
type SearchByNicknameRequest struct {
	//никнейм - уникальный
	Nickname string
}

//ответ SearchByNickname- поиск пользователей по никнейму
type SearchByNicknameResponse struct {
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
}
