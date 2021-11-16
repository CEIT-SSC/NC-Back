package modules

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/pkg/errors"
)

//room names
var roomNames = []string{"abc", "rwp", "def", "por", "ber", "erj"}

type RoomModule struct {
	RoomRepo repository.RoomRepository
}

func (r *RoomModule) CreateRoomForNewUser(userId int) error {
	for _, roomName := range roomNames {
		room := &models.Room{}
		room.RoomTitle = roomName
		room.UserID = userId
		room.RoomState = ""
		//TODO create room
		err := r.RoomRepo.CreateRoom(context.Background(), room)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
