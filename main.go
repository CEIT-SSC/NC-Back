package main

import (
	"github.com/ceit-ssc/nc_backend/api"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

func main() {

	dbConn := repository.CreateConnection()
	userRepo := repository.NewUserRepo(dbConn)
	roomRepo := repository.NewRoomRepo(dbConn)

	userModule := &modules.UserModule{UserRepo: userRepo}
	roomModule := &modules.RoomModule{RoomRepo: roomRepo}

	server := api.NewServer(userModule, roomModule)
	server.StartServer()
}
