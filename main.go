package main

import (
	"fmt"
	"github.com/ceit-ssc/nc_backend/api"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
)

func main() {
	dbConn, err := repository.CreateConnection()
	if err != nil {
		panic(err)
	}
	fmt.Println("server")
	userRepo := repository.NewUserRepo(dbConn)
	roomRepo := repository.NewRoomRepo(dbConn)
	tokenRepo := repository.NewTokenRepo(dbConn)

	userModule := &modules.UserModule{UserRepo: userRepo}
	roomModule := &modules.RoomModule{RoomRepo: roomRepo}
	tokenModule := &modules.TokenModule{TokenRepo: tokenRepo}

	server := api.NewServer(userModule, roomModule, tokenRepo, tokenModule)
	server.SetupRoutes()

	server.StartServer()
}
