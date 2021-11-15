package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/pkg/errors"
)

type UserRepository interface {
	CreateUser(ctx context.Context, newUser *models.User) error
	UpdateUserByField(ctx context.Context, user *models.User, fieldName string, value interface{}) error
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByStudentNumber(ctx context.Context, studentNumber int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	ExistsByUsernameAndPassword(ctx context.Context, user *models.User) (bool, error)
	RegisterUser(ctx context.Context, user *models.User) error
}

type UserRepoImpl struct {
	db *sql.DB
}

func (u UserRepoImpl) CreateUser(ctx context.Context, newUser *models.User) error {
	sqlStatement := `INSERT INTO users (username, password, student_number)
	VALUES ($1, $2, $3);`
	_, err := u.db.Exec(sqlStatement, newUser.Username, newUser.Password, newUser.StudentNumber)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u UserRepoImpl) UpdateUserByField(ctx context.Context, user *models.User, fieldName string, value interface{}) error {

	_, err := u.db.Exec("UPDATE users SET "+fieldName+" = $1 WHERE id = $2", value, user.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u UserRepoImpl) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	user := &models.User{}
	sqlStatement := `SELECT username, password, student_number FROM users WHERE id=$1;`
	err := u.db.QueryRow(sqlStatement, userID).Scan(user.Username, user.Password, user.StudentNumber)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (u UserRepoImpl) GetUserByStudentNumber(ctx context.Context, studentNumber int) (*models.User, error) {
	user := &models.User{}
	sqlStatement := `SELECT * FROM users WHERE student_number=$1;`
	row := u.db.QueryRow(sqlStatement, studentNumber)
	err := row.Scan(user.Username, user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (u UserRepoImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	sqlStatement := `SELECT * FROM users WHERE username=$1;`
	row := u.db.QueryRow(sqlStatement, username)
	err := row.Scan(user.ID, user.Username, user.Password, user.StudentNumber, user.Tokens)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (u UserRepoImpl) ExistsByUsernameAndPassword(ctx context.Context, user *models.User) (bool, error) {
	sqlStatement := `SELECT username , password FROM users WHERE username = $1 and password = $2;`

	_, err := u.db.Exec(sqlStatement, user.Username, user.Password)
	if err == sql.ErrNoRows{
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u UserRepoImpl) RegisterUser(ctx context.Context, user *models.User) error {
	sqlStatement := `SELECT username , password FROM users WHERE username = $1 and password = $2;`
	var username string
	var password string
	row := u.db.QueryRow(sqlStatement, user.Username, user.Password)
	switch err := row.Scan(&username, &password); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return errors.New("User is already registered")
	default:
		fmt.Println(err)
	}
	return nil
}

func NewUserRepo(dbConn *sql.DB) UserRepository {
	return &UserRepoImpl{db: dbConn}
}
