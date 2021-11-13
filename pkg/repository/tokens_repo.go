package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/pkg/errors"
)

type UserTokens interface {
	CreateNewToken(ctx context.Context, userID int, token string) error
	GetUserTokens(ctx context.Context, userID int) ([]string, error)
}

func CreateNewToken(ctx context.Context, userID int, token string) error {
	db := models.CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO user_tokens (user_id, token)
	VALUES ($1, $2);`
	err := db.QueryRow(sqlStatement, userID, token)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(err)
	return nil
}

func GetUserTokens(ctx context.Context, userID int) ([]string, error) {
	db := models.CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM user_tokens WHERE user_id = $1;`
	row := db.QueryRow(sqlStatement, userID)
	err := row.Scan("token")
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned")
		return Tokens{}, errors.WithStack(err)
	}
	if err != nil {
		fmt.Println(err)
		return Tokens{}, errors.WithStack(err)
	}
	return Tokens{}, nil
	return nil, nil
}

