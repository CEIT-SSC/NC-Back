package controllers

import (
	"context"
	"fmt"
	"github.com/ceit-ssc/nc_backend/internal/modules"
	error2 "github.com/ceit-ssc/nc_backend/pkg/error"
	"github.com/ceit-ssc/nc_backend/pkg/models"
	"github.com/ceit-ssc/nc_backend/pkg/repository"
	"github.com/ceit-ssc/nc_backend/pkg/token"
	"github.com/gin-gonic/gin"
	_ "github.com/pkg/errors"
	"math/rand"
)

func RegisterController(userModule *modules.UserModule, roomModule *modules.RoomModule) gin.HandlerFunc {

	return func(context *gin.Context) {

		user := &models.User{}
		err := context.ShouldBindJSON(user)
		if err != nil {
			context.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		userId, err := userModule.RegisterNewUser(user)
		if err == error2.ErrUserIsRegistered {
			context.JSON(400, gin.H{
				"type": "user_error",
				"error": err.Error(),
			})
			return
		}
		if err != nil {
			context.JSON(500, gin.H{
				"type": "user_error",
				"error": err.Error(),
			})
			return
		}

		err = roomModule.CreateRoomForNewUser(userId)
		if err != nil {
			userDeleteErr := userModule.DeleteUserByID(userId)
			context.JSON(500, gin.H{
				"type": "room_error",
				"error": err.Error(),
				"user_delete_error": userDeleteErr,
			})
			return
		}

		context.JSON(201, gin.H{
			"message": "user created successfully",
		})
	}
}

func LoginController(userModule *modules.UserModule, tokenRepo repository.UserTokens) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		user := &models.LoginUser{}
		err := ctx.ShouldBindJSON(user)
		if err != nil {
			ctx.JSON(422, gin.H{
				"error":   err.Error(),
				"message": "invalid request body",
			})
			return
		}
		fmt.Println(user)
		userInfo,err := userModule.GetUserByUsername(user.Username)
		if userInfo == nil && err == error2.ErrNoUserFound{
			ctx.JSON(404,gin.H{
				"error":  err.Error(),
				"message": "no user found",
			})
		}
		if err != nil{
			ctx.JSON(500,gin.H{
				"error":  err.Error(),
			})
		}
		fmt.Println(userInfo)
		UserToken, err := token.NewToken(context.Background(), fmt.Sprintf("%d", userInfo.ID), rand.Int())
		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = tokenRepo.CreateNewToken(context.Background(), userInfo.ID, UserToken)
		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
				"type": "token_repo",
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

func LogoutController(tokenRepo repository.UserTokens) gin.HandlerFunc {

	return func(ctx *gin.Context) {
	}

}
