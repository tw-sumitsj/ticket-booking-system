package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ping Successful",
	})
}
