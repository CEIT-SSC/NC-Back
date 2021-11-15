package error

import "github.com/pkg/errors"

var (
	ErrTokenMissing  = errors.New("token is missing")
	ErrInvalidToken  = errors.New("invalid token")
	ErrNotRegistered = errors.New("invalid User id")
	ErrUserIsRegiseterd = errors.New("User is Registered before")
)
