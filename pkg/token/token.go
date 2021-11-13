package token

import (
	"context"
	"github.com/dgrijalva/jwt-go"
)
func NewToken(ctx context.Context, input string) (string, error) {
	claims := jwt.MapClaims{"usere_id":input}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte("key"))
	if err != nil {
		return t, err
	}
	return t, nil
}