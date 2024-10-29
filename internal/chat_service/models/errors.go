package models

import "errors"

var (
	ErrNotFoundNickname= errors.New("nickname not found")
	ErrNotFoundFriend= errors.New("friend nickname not found")
	ErrGetMessageFailed = errors.New("failed get message")
	ErrWriteFriendFailed = errors.New("failed to save a friend's message")
)