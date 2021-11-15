package modules

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/ceit-ssc/nc_backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type UserModule struct {
	UserRepo repository.UserRepository
}

// TODO : implement login method
func (u *UserModule) LoginUser(c *gin.Context, ctx context.Context, user *models.User) (int, error) {
	user, err := repository.UserRepoImpl{}.LoginUser(ctx, user)
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
func (u *UserModule) RegisterNewUser(ctx context.Context, user models.User) (int, error) {
	err := repository.UserRepoImpl{}.RegisterUser(ctx, &user)
	if err != nil {
		return 0, err
	}
	err = repository.UserRepoImpl{}.CreateUser(ctx, &user)
	if err != nil {
		return 0, err
	}
	return -1, nil
}
