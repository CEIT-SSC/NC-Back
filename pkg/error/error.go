package error

import "github.com/pkg/errors"

var (
	ErrTokenMissing  = errors.New("token is missing")
	ErrInvalidToken  = errors.New("invalid token")
	ErrNotRegistered    = errors.New("invalid User id")
	ErrUserIsRegistered = errors.New("User is Registered before")
	ErrWrongPass = errors.New("Username and password don't match")
	ErrNoUserFound = errors.New("no user found")
)
