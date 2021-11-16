package api

import (
	"github.com/ceit-ssc/nc_backend/api/controllers"
	"github.com/ceit-ssc/nc_backend/api/middlewares"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//implement router and their handlers here

// USE GIN FRAMEWORK FOR GOD'S SAKES

// initialize modules in here and pass them and use them in controllers
type Server struct {
	UserModule *modules.UserModule
	RoomModule *modules.RoomModule
	TokenRepo repository.UserTokens
	router *gin.Engine
}

func NewServer(userModule *modules.UserModule, roomModule *modules.RoomModule, tokenRepo repository.UserTokens) *Server {
	router := gin.Default()
	return &Server{
		router: router,
		UserModule: userModule,
		RoomModule: roomModule,
		TokenRepo: tokenRepo,
	}
}

func (s *Server) StartServer() {
	log.Fatal(http.ListenAndServe(":8080", s.router))
}

//TODO: add these routes
//     /user/register: should create user (using modules) and their empty rooms
//     /user/login: should return token

func (s *Server) SetupRoutes() {

	s.router.POST("/user/register", controllers.RegisterController(s.UserModule, s.RoomModule))
	s.router.POST("/user/login", controllers.LoginController(s.UserModule, s.TokenRepo))
	s.router.POST("/user/logout", middlewares.IsAuthenticated(s.TokenRepo), controllers.LogoutController(s.TokenRepo))
}
