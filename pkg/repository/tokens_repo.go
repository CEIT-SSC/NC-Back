package repository

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type UserTokens interface {
	CreateNewToken(ctx context.Context, userID int, token string) error
	GetUserTokens(ctx context.Context, userID int) ([]string, error)
	RemoveToken(ctx context.Context, token string, id int) error
	CheckUserId(ctx context.Context, id string) bool
}

type UserTokenImpl struct {
	db *sql.DB
}

func NewTokenRepo(db *sql.DB) UserTokens {
	return &UserTokenImpl{db: db}
}

func (u *UserTokenImpl) GetUserTokens(ctx context.Context, userID int) ([]string, error) {
	var tokens []string

	sqlStatement := `SELECT token FROM user_tokens WHERE user_id = $1;`
	rows, err := u.db.Query(sqlStatement, userID)
	if err == sql.ErrNoRows {
		return nil, errors.New("no token found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for rows.Next() {
		var token string
		err := rows.Scan(&token)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (u UserTokenImpl) CreateNewToken(ctx context.Context, userID int, token string) error {
	sqlStatement := `INSERT INTO user_tokens (user_id, token)
	VALUES ($1, $2);`
	_, err := u.db.Exec(sqlStatement, userID, token)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u UserTokenImpl) RemoveToken(ctx context.Context, token string, id int) error {
	_, err := u.db.Exec("DELETE FROM user_tokens WHERE token = $1 and user_id = $2", token, id)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u UserTokenImpl) CheckUserId(ctx context.Context, id string) bool {
	row := u.db.QueryRow("SELECT * FROM user_tokens WHERE user_id= $1", id)
	var token string
	err := row.Scan(&token)
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
