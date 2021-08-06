package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	//TODO: Read user credentials. Validate. Create session. Set session.
	c.Set("SESSION", "SECRET@SESSION@ID")
}
