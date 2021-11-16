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

//TODO: DO me
func (u *UserModule) RegisterNewUser(user *models.User) (int,error) {

	exists, err := u.UserRepo.ExistsByUsernameAndPassword(context.Background(), user.Username, user.Password)
	if err != nil {
		return -1, err
	}
	if exists {
		return -1, error2.ErrUserIsRegistered
	}
	id, err := u.UserRepo.CreateUser(context.Background(), user)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (u *UserModule) GetUserByUsername(username string)  (*models.User,error){
	user, err := u.UserRepo.GetUserByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (u *UserModule) DeleteUserByID(userID int) error {
	return u.UserRepo.DeleteUserByID(context.Background(), userID)
}