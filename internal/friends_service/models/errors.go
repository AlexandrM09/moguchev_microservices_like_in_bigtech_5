package models

import "errors"

var (
	ErrNotFoundNickname= errors.New("nickname not found")
	ErrNotFoundFriend= errors.New("friend nickname not found")
	ErrAddFailed = errors.New("couldn't add a friend")
	ErrRemoveFailed = errors.New("couldn't remuve a friend")
	ErrConfirmFailed = errors.New("the friendship request could not be confirmed or rejected")
	ErrGetListFailed = errors.New("failed get list")
)
