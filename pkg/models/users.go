package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type User struct {
	ID            int      `json:"-"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	StudentNumber int      `json:"student_number"`
	Tokens        []string `json:"-"`
}

func CreateUser(ctx context.Context, newUser User) (User, error) {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO users (username, password, student_number)
	VALUES ($1, $2, $3);`
	err := db.QueryRow(sqlStatement, newUser.ID, newUser.Password, newUser.Password)
	if err != nil {
		return newUser, errors.WithStack(err)
	}
	return newUser, nil
}

func GetUserByID(ctx context.Context, userID int) (User, error) {
	user := &User{}
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	row := db.QueryRow(sqlStatement, userID)
	err := row.Scan(user.Username, user.Password, user.StudentNumber)
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned")
		return User{}, errors.WithStack(err)
	}
	if err != nil {
		fmt.Println(err)
		return User{}, errors.WithStack(err)
	}
	return User{}, nil
}

func GetUserByStudentNumber(ctx context.Context, studentNumber int) (User, error) {
	user := &User{}
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM users WHERE student_number=$1;`
	row := db.QueryRow(sqlStatement, studentNumber)
	err := row.Scan(user.Username, user.Password)
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned")
		return User{}, errors.WithStack(err)
	}
	if err != nil {
		fmt.Println(err)
		return User{}, errors.WithStack(err)
	}
	return User{}, nil
}

func UpdateUserByField(ctx context.Context, user User, fieldName string, value interface{}) error {
	db := CreateConnection()
	defer db.Close()
	err, _ := db.Exec("UPDATE users SET "+fieldName+" = $1 WHERE id = $2", value, user.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

