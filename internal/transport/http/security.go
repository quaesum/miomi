package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AdminTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			fmt.Println("AHTUNG!")
		}
	}
}
