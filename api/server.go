package api

import (
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

//router.POST("/user/new", controllers.Register)
//router.POST("/user/login", app.AuthorizeJWT, controllers.Authenticate)
//router.POST("/user/upload", controllers.Upload)
//router.GET("/user/download", app.AuthorizeJWT, controllers.Download)
//router.POST("/user/give/permission", controllers.GivePermission)
func setupRoutes() {
	router := gin.Default()
	router.POST("/user/register", func(context *gin.Context) {
		users := models.User{}
		err := context.ShouldBindJSON(&users)
		if err != nil {
			context.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}
		module := modules.UserModule{}
		_, err = module.RegisterNewUser(context, users)
		if err != nil {
			errors.WithStack(err)
		}
	})
	router.GET("user/login", func(context *gin.Context) {

	})

}
