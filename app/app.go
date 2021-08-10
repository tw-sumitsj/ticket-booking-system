package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	registerRoutes(router)
	router.Run()
	fmt.Println("Server Started")
}
