package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.Status(http.StatusOK)
	}
}
