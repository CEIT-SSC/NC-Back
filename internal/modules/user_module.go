package modules

import (
	"context"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/ceit-ssc/nc_backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserModule struct {
	UserRepo repository.UserRepository
}

// TODO : implement login method
func (u *UserModule) LoginUser(c *gin.Context, ctx context.Context, user *models.User) (int, error) {
	user, err := repository.UserRepoImpl{}.LoginUser(ctx, user) //wrong
	if err != nil {
		return 0, err
	}
	loginToken, _ := token.NewToken(c, string(rune(user.ID)), 1)
	c.JSON(200, gin.H{
		"message":        "user is logged in",
		"username":       user.Username,
		"student_number": user.StudentNumber,
		"loginToken":     loginToken,
	})

	return 0, nil
}

//TODO: DO me
func (u *UserModule) RegisterNewUser(user models.User) (error) {

	exists, err := u.UserRepo.ExistsByUsernameAndPassword(context.Background(), &user)
	if err != nil {
		return err
	}
	if exists {
		return error2.ErrUserIsRegiseterd
	}
	err = u.UserRepo.CreateUser(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}
