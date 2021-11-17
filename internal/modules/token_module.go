package modules

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

type TokenModule struct {
	TokenRepo repository.UserTokens
}

func (t *TokenModule) RemoveToken(ctx context.Context, token string, id int) error {
	err := t.TokenRepo.RemoveToken(ctx, token, id)
	if err != nil {
		return err
	}
	return nil
}
