package modules

import (
	"context"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

type UserModule struct {
	UserRepo repository.UserRepository
}

// TODO : implement login method
func (u *UserModule) LoginUser(ctx context.Context, user *models.User) error {
	exists, err := u.UserRepo.ExistsByUsernameAndPassword(ctx, user)
	if err != nil {
		return err
	}
	if !exists {
		return error2.ErrWrongPass
	}

	return nil
}

//TODO: DO me
func (u *UserModule) RegisterNewUser(user models.User) error {

	exists, err := u.UserRepo.ExistsByUsernameAndPassword(context.Background(), &user)
	if err != nil {
		return err
	}
	if exists {
		return error2.ErrUserIsRegistered
	}
	err = u.UserRepo.CreateUser(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModule) GetUserByUsername(username string)  *models.User{
	user, err := u.UserRepo.GetUserByUsername(context.Background(), username)
	if err != nil {
		return nil
	}
	return user
}


