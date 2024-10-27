package usecase

import (
	"context"

	"github.com/AlexandrM09/moguchev_microservices_like_in_bigtech_5/internal/users_service/models"
)

// Usecase - port (первичный)
//
// ...
type Usecase interface {
	//SignUpUser - регистрация пользователя (по почте и паролю)
	SignUpUserEmail(ctx context.Context, req models.SignUpUserEmailRequest) (*models.SignUpUserEmailResponse, error)
	//SignInUser - вход/авторизация (по почте и паролю)
	SignInUserEmail(ctx context.Context, req models.SignInUserPostRequest) (*models.SignInUserPostResponse, error)
	//SignInUserOauth - вход/авторизация (Oauth), регистрация если нет такого пользователя
	SignInUserOauth(ctx context.Context, req models.SignInUserOauthRequest) (*models.SignInUserOauthResponse, error)
	//UpdateProfile - редактирование профиля пользователя (никнейм - уникальный, информация о себе, аватарка)
	UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) (*models.UpdateProfileResponse, error)
	//SearchByNickname- поиск пользователей по никнейму
	SearchByNickname(ctx context.Context, req models.SearchByNicknameRequest) (*models.SearchByNicknameResponse, error)
}

type (
	// ChatRepository - port (вторичный)
	UsersRepository interface {
		//SignUpUser - регистрация пользователя (по почте и паролю)
		SignUpUserEmail(ctx context.Context, req models.SignUpUserEmailRequest) (*models.SignUpUserEmailResponse, error)
		//SignInUser - вход/авторизация (по почте и паролю)
		SignInUserEmail(ctx context.Context, req models.SignInUserPostRequest) (*models.SignInUserPostResponse, error)
		//SignInUserOauth - вход/авторизация (Oauth), регистрация если нет такого пользователя
		SignInUserOauth(ctx context.Context, req models.SignInUserOauthRequest) (*models.SignInUserOauthResponse, error)
		//UpdateProfile - редактирование профиля пользователя (никнейм - уникальный, информация о себе, аватарка)
		UpdateProfile(ctx context.Context, req models.UpdateProfileRequest) (*models.UpdateProfileResponse, error)
		//SearchByNickname- поиск пользователей по никнейму
		SearchByNickname(ctx context.Context, req models.SearchByNicknameRequest) (*models.SearchByNicknameResponse, error)
	}
)

// Deps -
type Deps struct {
	// Adapters
	UsersRepository UsersRepository
}

var _ Usecase = (*usecase)(nil)

type usecase struct {
	Deps
}

// бизнес сервис
func NewUsecase(d Deps) *usecase {
	return &usecase{Deps: d}
}
