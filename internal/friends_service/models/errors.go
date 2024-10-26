package models

import "errors"

var (
	ErrNotFoundFriend= errors.New("friend not found")
	ErrGetMessageFailed = errors.New("failed to save a friend's message")
	ErrWriteFriendFailed = errors.New("failed to save a friend's message")
)