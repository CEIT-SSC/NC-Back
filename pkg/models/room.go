package models

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type Room struct {
	UserID     int    `json:"-"`
	RoomTitle  string `json:"room_title"`
	RoomState  string `json:"room_state"`
	IsComplete bool   `json:"is_complete"`
	Score      int    `json:"score"`
}

func CreateRoom(context context.Context, newRoom Room) error {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO rooms (room_title, room_state, is_complete, score)
	VALUES ($1, $2, $3, $4)`
	err := db.QueryRow(sqlStatement, newRoom.RoomTitle, newRoom.RoomState, newRoom.IsComplete, newRoom.Score)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(err)
}

func UpdateRoomByField(ctx context.Context, roomTitle string, userID int, fieldName string, value interface{}) error {
	db := CreateConnection()
	defer db.Close()
	err, _ := db.Exec("UPDATE users SET "+fieldName+" = $1 WHERE (ID = $2 AND room_title = $3) ", value, userID, roomTitle)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func DeleteRoom(ctx context.Context, title string, userID int) error {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `
	DELETE FROM users
	WHERE (user_id = $1 AND room_title = $2);`
	err, _ := db.Exec(sqlStatement, userID, title)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func GetRoomByUser(ctx context.Context, title string, userID int) (Room, error) {
	room := &Room{}
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM rooms WHERE (user_id = $1 and room_title = $2);`
	row := db.QueryRow(sqlStatement, userID, title)
	err := row.Scan(room.RoomState, room.IsComplete, room.Score)
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned")
		return Room{}, errors.WithStack(err)
	}
	if err != nil {
		fmt.Println(err)
		return Room{}, errors.WithStack(err)
	}
	return Room{}, nil

}

