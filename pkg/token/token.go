package token

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type authCustomClaims struct {
	input string `json:"name"`
	jwt.StandardClaims
}

func NewToken(ctx context.Context, input string) (string, error) {
	claims := jwt.MapClaims{"usere_id":&authCustomClaims{
		input,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 90).Unix(),
			//Issuer:    service.issure,
			IssuedAt: time.Now().Unix(),
		},
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte("key"))
	if err != nil {
		return t, err
	}
	return t, nil
}