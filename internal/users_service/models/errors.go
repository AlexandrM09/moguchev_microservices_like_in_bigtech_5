package models

import "errors"

var (
	ErrSignUpUserFailed = errors.New("sign up user failed")
	ErrSignUpUserFailedUserExist = errors.New("sign up user failed - user exist")
	ErrSignUpUserFailedEmail  = errors.New("sign up(email) user failed")
	ErrSignUpUserFailedEmailUserExist = errors.New("sign up(email) user failed - user exist")
	ErrSignInUserEmailFailed = errors.New("sign in user(email) failed")
	ErrSignInUserOauthFailed = errors.New("sign in user(oauth) failed")
	ErrUpdateProfileFailed = errors.New("update profile failed")
	ErrSearchByNicknameFailed = errors.New("search by nickname failed")	
)