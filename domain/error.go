package domain

import "errors"

var (
    ErrUserNotFound    = errors.New("user not found")
    ErrEmailAlreadyUsed = errors.New("email already used")
    ErrInvalidToken    = errors.New("invalid token")
    ErrTokenExpired    = errors.New("token expired")
)
