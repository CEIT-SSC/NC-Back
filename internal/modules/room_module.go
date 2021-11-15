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

func (u *UserModule) RegisterNewRoom(ctx context.Context, user models.User) (int, error) {
	for _, roomName := range roomNames {
		room := models.Room{}
		room.RoomTitle = roomName
		room.UserID = user.ID
		err := repository.RoomRepoImpl{}.CreateRoom(ctx, &room)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}
	return -1, nil
}
