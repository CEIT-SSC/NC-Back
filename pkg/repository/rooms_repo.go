package repository

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
)

type RoomRepository interface {
	CreateRoom(context.Context, *models.Room) error
	UpdateRoomByField(ctx context.Context, roomTitle string, userID int, fieldName string, value interface{})error
	DeleteRoom(ctx context.Context, title string, userID int) error
	GetRoomByUser(ctx context.Context, title string, userID int) (*models.Room, error)
}