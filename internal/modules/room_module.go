package modules

import (
	"context"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

type RoomModule struct {
	RoomRepo repository.RoomRepository
}

func (u *UserModule) RegisterNewRoom(ctx context.Context, user models.User) (int, error) {

	//repository.RoomRepoImpl{}.CreateRoom()

	return -1, nil
}
