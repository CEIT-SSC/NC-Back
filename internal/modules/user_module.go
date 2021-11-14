package modules

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

type UserModule struct {
	UserRepo repository.UserRepository
}

//TODO: DO me
func (u *UserModule) RegisterNewUser(ctx context.Context, user models.User) (int, error) {

	err := repository.UserRepoImpl{}.CreateUser(ctx, &user)
	if err != nil {
		return 0, err
	}

	return -1, nil
}
