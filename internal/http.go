package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"madmax/internal/mysql"
	"net/http"
	"strconv"
	"time"
)

func HTTPHandler(router *gin.Engine) {
	router.GET("/", basicInfoHandler)

	userGroup := router.Group("/user/v1")
	userGroup.POST("/signup")
	userGroup.GET("/:id", getUserByIDHandler)
	//user.GET("/animal", animalsInfo)
}

/*
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

*/

func getUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	uID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*2)
	userByID(tctx, uID)
}

func basicInfoHandler(c *gin.Context) {
	userinfo, err := mysql.GetUserBasicInfo(context.Background(), 1)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userinfo)
	send := fmt.Sprintf("%+v", userinfo)

	c.String(http.StatusOK, send)

}

func animalsInfo(c *gin.Context) {
	animalinfo1, err := mysql.GetAnimalBasicInfo(context.Background(), 1)
	animalinfo2, err := mysql.GetAnimalBasicInfo(context.Background(), 2)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	animal1 := fmt.Sprintf("%+v", animalinfo1)
	animal2 := fmt.Sprintf("%+v", animalinfo2)
	c.String(http.StatusOK, animal1)
	c.String(http.StatusOK, animal2)
}
