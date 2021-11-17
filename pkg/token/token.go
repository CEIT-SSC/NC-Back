package token

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strconv"
)

func NewToken(ctx context.Context, input string, randomNumber int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": input,
		"rand":    randomNumber,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte("key"))
	if err != nil {
		return t, err
	}

	return t, nil
}

func ParseToken(ctx context.Context, token2 string) (*jwt.Token, error) {
	token, err := jwt.Parse(token2, func(token *jwt.Token) (interface{}, error) {
		return []byte("key"), nil
	})
	return token, err
}

func GetUserID(ctx context.Context, token string) (int, error) {
	parsedToken, err := ParseToken(ctx, token)
	if err != nil {
		return -1, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("Failed to parse token")
	}
	userIDInt, _ := strconv.Atoi(claims["user_id"].(string))
	return userIDInt, nil
}
