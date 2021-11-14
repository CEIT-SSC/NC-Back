package api

import (
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/gin-gonic/gin"
)

//implement router and their handlers here

// USE GIN FRAMEWORK FOR GOD'S SAKES

// initialize modules in here and pass them and use them in controllers
type Server struct {
	UserModule *modules.UserModule
	RoomModule *modules.RoomModule
}

func NewServer(userModule *modules.UserModule, roomModule *modules.RoomModule) *Server {
	return &Server{
		UserModule: userModule,
		RoomModule: roomModule,
	}
}

func (s *Server) StartServer() {

}

//TODO: add these routes
//     /user/register: should create user (using modules) and their empty rooms
//     /user/login: should return token

func setupRoutes() {
	router := gin.Default()
	router.GET("/user/login", func(context *gin.Context) {
		user := &models.User{}
		context.BindJSON(user)
		repository.UserRepoImpl{}.Login(context, user)
	})
}
