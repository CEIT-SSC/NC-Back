package repository

import "context"

type UserTokens interface {
	CreateNewToken(ctx context.Context, userID int, token string) error
	GetUserTokens(ctx context.Context, userID int) ([]string, error)
}