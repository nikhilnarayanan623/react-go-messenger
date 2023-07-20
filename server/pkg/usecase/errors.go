package usecase

import "errors"

var (
	// login
	ErrEmptyLoginCredentials = errors.New("all login credentials are empty")
	ErrUserNotExist          = errors.New("user not exist with given login credentials")
	ErrWrongPassword         = errors.New("password doesn't match")

	// refresh token
	ErrInvalidRefreshToken    = errors.New("invalid refresh token")
	ErrRefreshSessionNotExist = errors.New("there is no refresh token session for this token")
	ErrRefreshSessionExpired  = errors.New("refresh token expired in session")
	ErrRefreshSessionBlocked  = errors.New("refresh token blocked in session")

	// signup
	ErrUserAlreadyExit = errors.New("user already exist")
)
