package middlewares

import (
	"context"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/ceit-ssc/nc_backend/pkg/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func IsAuthenticated (tokenRepo repository.UserTokens) gin.HandlerFunc{
	return func (c *gin.Context){
		tokenHeader := c.GetHeader("Authorization")
		userToken,_ := GetToken(tokenHeader)
		userID, err := GetUserID(userToken)

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
		tokens, err := tokenRepo.GetUserTokens(context.Background(), userID)

		tokenExists := tokenExistsOnList(tokens, userToken)
		if !tokenExists{
			c.JSON(403, gin.H{
				"error": "user is not authenticated",
			})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}

func GetToken(tokenHeader string) (string, error){
	if tokenHeader == "" {
		return "", error2.ErrTokenMissing
	}
	//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return "", error2.ErrInvalidToken
	}
	//Grab the token part
	return splitted[1], nil
}

func GetUserID(tokenPart string) (int, error) {
	userID, err := token.GetUserID(context.Background(), tokenPart)
	if err != nil {
		return -1, error2.ErrInvalidToken
	}
	return userID, nil
}
func tokenExistsOnList(tokens []string, userToken string) bool {
	for _, token2:= range tokens{
		if token2 == userToken{
			return true
		}
	}
	return false
}