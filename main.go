package main

import (
	"github.com/ceit-ssc/nc_backend/api"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/models"
)

func main() {
	models.CreateConnection()
	userModule := &modules.UserModule{}
	roomModule := &modules.RoomModule{}

	server := api.NewServer(userModule, roomModule)
	server.StartServer()
}
