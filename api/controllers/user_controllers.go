package controllers

import (
	"context"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/gin-gonic/gin"
	_ "github.com/pkg/errors"
)

func RegisterController(module *modules.UserModule) gin.HandlerFunc {
	return func(context *gin.Context) {
		user := models.User{}
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		err = module.RegisterNewUser(user)

		if err == error2.ErrUserIsRegiseterd {
			context.JSON(400,gin.H{
				"error": err.Error(),
			})
			return
		}
		if err != nil {
			context.JSON(500,gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(201, gin.H{
			"message": "user created successfully",
		})
	}
}

func LoginController( userModule *modules.UserModule) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO: step1: query to userRepo using userModule to get ID
		 //TODO: step2: Create New token using userId




		userLogin := models.User{}
		err := ctx.ShouldBindJSON(&userLogin)
		if err != nil {
			ctx.JSON(422, gin.H{
				"error":   true,
				"message": "invalid request body",
			})
			return
		}

		_, err = modules.LoginUser(ctx, &models.User{})//wrong
		if err != nil {
			ctx.JSON(422, gin.H{
				"error":   true,
				"message": "cannot login user",
			})
		}
	}
}
