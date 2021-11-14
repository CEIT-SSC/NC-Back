package middlewares

import (
	"context"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/ceit-ssc/nc_backend/pkg/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func IsAuthenticated(c *gin.Context, tokenRepo repository.UserTokens) {
	tokenHeader := c.GetHeader("Authorization")
	user_id, err := GetUserID(tokenHeader)

	switch err {
	case error2.ErrTokenMissing:
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	case error2.ErrInvalidToken:
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}
	//TODO: Check if user id is present in token table and then if it was present pass it to next

	c.Set("user", user_id)
	r := repository.UserTokenImpl{}.CheckUserId(c, user_id)
	if !r {
		c.JSON(403, gin.H{
			"error": error2.ErrNotRegistered,
		})
	}
	c.Next()
}

func GetUserID(tokenHeader string) (string, error) {
	if tokenHeader == "" {
		return "", error2.ErrTokenMissing
	}
	//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return "", error2.ErrInvalidToken
	}
	//Grab the token part
	tokenPart := splitted[1]
	user_id, err := token.GetUserID(context.Background(), tokenPart)
	if err != nil {
		return "", error2.ErrInvalidToken
	}
	return user_id, nil
}
