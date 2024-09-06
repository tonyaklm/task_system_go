package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task_system_go/token"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
		c.Abort()
		return
	}
	claims, err := token.ValidateToken(clientToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Set("user_id", claims.UserID)
	c.Next()
}
