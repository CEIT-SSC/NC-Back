package main

import (
	"github.com/ceit-ssc/nc_backend/api"
	"github.com/ceit-ssc/nc_backend/internal/modules"
)

func main(){
	userModule := &modules.UserModule{}
	roomModule := &modules.RoomModule{}

	server := api.NewServer(userModule,roomModule)
	server.StartServer()
}



