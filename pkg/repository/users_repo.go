package repository

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, newUser *models.User) error
	UpdateUserByField(ctx context.Context, user *models.User, fieldName string, value interface{})error
	GetUserByID(ctx context.Context, userID int) (*models.User,error)
	GetUserByStudentNumber (ctx context.Context, studentNumber int) (*models.User, error)
}