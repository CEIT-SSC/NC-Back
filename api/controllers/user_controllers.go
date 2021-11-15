package controllers

import (
	"github.com/ceit-ssc/nc_backend/internal/modules"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func RegisterController(module *modules.UserModule) gin.HandlerFunc {
	return func(context *gin.Context) {
		users := models.User{}
		err := context.ShouldBindJSON(&users)
		if err != nil {
			context.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		_, err = module.RegisterNewUser(context, users)
		if err != nil {
			errors.WithStack(err)
		}
	}
}

func LoginController(c *gin.Context, modules *modules.UserModule) gin.HandlerFunc {
	return func(context *gin.Context) {
		userLogin := models.User{}
		err := context.ShouldBindJSON(&userLogin)
		if err != nil {
			context.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		_, err = modules.LoginUser(c, context, &models.User{})
		if err != nil {
			context.JSON(422, gin.H{
				"error":   true,
				"message": "cannot login user",
			})
		}
	}
}
