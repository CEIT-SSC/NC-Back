package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/pkg/errors"
)

type RoomRepository interface {
	CreateRoom(context.Context, *models.Room) error
	UpdateRoomByField(ctx context.Context, roomTitle string, userID int, fieldName string, value interface{})error
	DeleteRoom(ctx context.Context, title string, userID int) error
	GetRoomByUser(ctx context.Context, title string, userID int) (*models.Room, error)
}

type RoomRepoImpl struct{
	db *sql.DB
}

func (r RoomRepoImpl) CreateRoom(ctx context.Context, room *models.Room) error {

	sqlStatement := `INSERT INTO rooms (room_title, room_state, is_complete, score)
	VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(sqlStatement, room.RoomTitle, room.RoomState, room.IsComplete, room.Score)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r RoomRepoImpl) UpdateRoomByField(ctx context.Context, roomTitle string, userID int, fieldName string, value interface{}) error {
	_, err := r.db.Exec("UPDATE users SET "+fieldName+" = $1 WHERE (ID = $2 AND room_title = $3) ", value, userID, roomTitle)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r RoomRepoImpl) DeleteRoom(ctx context.Context, title string, userID int) error {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM users WHERE (user_id = $1 AND room_title = $2);`
	_, err := db.Exec(sqlStatement, userID, title)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r RoomRepoImpl) GetRoomByUser(ctx context.Context, title string, userID int) (*models.Room, error) {
	room := &models.Room{
		UserID: userID,
		RoomTitle: title,
	}
	sqlStatement := `SELECT room_state, is_complete, score FROM rooms WHERE user_id = $1 and room_title = $2;`
	err := r.db.QueryRow(sqlStatement, userID, title).Scan(room.RoomState, room.IsComplete, room.Score)
	if err == sql.ErrNoRows {
		return nil, errors.New("found no room")
	}
	if err != nil {
		fmt.Println(err)
		return nil, errors.WithStack(err)
	}
	return nil, nil
}

func NewRoomRepo (db *sql.DB) RoomRepository{
	return &RoomRepoImpl{db: db}
}