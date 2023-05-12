package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"madmax/internal/mysql"
	"net/http"
)

func HTTPHandler(router *gin.Engine) {
	router.GET("/", basicInfoHandler)
	user := router.Group("/user/v1")
	user.GET("/", basicInfoHandler)
}

func basicInfoHandler(c *gin.Context) {
	userinfo, err := mysql.GetUserBasicInfo(context.Background(), 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userinfo)

	c.String(http.StatusOK, "Welcome Gin Server")
}
