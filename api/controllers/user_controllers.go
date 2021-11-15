package controllers

import (
	"context"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/token"
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

		if err == error2.ErrUserIsRegistered {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(201, gin.H{
			"message": "user created successfully",
		})
	}
}

func LoginController(userModule *modules.UserModule) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO: step1: query to userRepo using userModule to get ID
		//TODO: step2: Create New token using userId
		user := models.User{}
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(422, gin.H{
				"error":   err.Error(),
				"message": "invalid request body",
			})
			return
		}
		err = userModule.LoginUser(ctx, &models.User{})
		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
		}

		userInfo := userModule.GetUserByUsername(user.Username)
		UserToken, err := token.NewToken(context.Background(), string(rune(userInfo.ID)), 3)
		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"username": user.Username,
			"student_number": userInfo.StudentNumber,
			"token": UserToken,
		})
	}

}
