package v1

import (
	"github.com/gin-gonic/gin"
	"madmax/internal/utils"
	"net/http"
)

func AdminTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := utils.GetUserRole(c)
		if err != nil || role != utils.UserRoleAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		return
	}
}

func UserTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := utils.GetUserRole(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		return
	}
}

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With, access-control-allow-origin")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
