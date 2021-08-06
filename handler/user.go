package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sujithps/ticket-booking-system/contract"
	"net/http"
)

func UserHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userRequest := contract.UserRequest{}
		err := ctx.ShouldBindBodyWith(&userRequest, binding.JSON)
		if err != nil {
			fmt.Printf("Got error %+v \n", err)
		}

		if sessionId, ok := ctx.Get("SESSION"); ok {
			fmt.Printf("Got Session ID %s \n", sessionId)
			ctx.JSON(http.StatusOK, contract.UserResponse{Authenticated: true})
			return
		}

		ctx.JSON(http.StatusUnauthorized, contract.UserResponse{Authenticated: false})

	}
}
