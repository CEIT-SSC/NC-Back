package api

import "github.com/ceit-ssc/nc_backend/internal/modules"

//implement router and their handlers here

// USE GIN FRAMEWORK FOR GOD'S SAKES

// initialize modules in here and pass them and use them in controllers
type Server struct {
	UserModule *modules.UserModule
	RoomModule *modules.RoomModule
}

func NewServer(userModule *modules.UserModule, roomModule *modules.RoomModule)*Server{
	return &Server{
		UserModule: userModule,
		RoomModule: roomModule,
	}
}

func (s *Server) StartServer(){

}



func setupRoutes(){

}