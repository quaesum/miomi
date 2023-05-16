package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"madmax/internal/entity"
	"madmax/internal/mysql"
	"net/http"
	"strconv"
	"time"
)

func HTTPHandler(router *gin.Engine) {
	router.GET("/", basicInfoHandler)

	userGroup := router.Group("/user/v1")
	userGroup.GET("/:id", getUserByIDHandler)
	userGroup.GET("/", getAllUsersHandler)
	userGroup.POST("/signup", userSignupHandler)
	userGroup.POST("/:id", updateUserHandler)

	animalGroup := router.Group("/animal/v1")
	animalGroup.GET("/:id", getAnimalByIDHandler)
	animalGroup.GET("/", getAllUsersDHandler)
	animalGroup.POST("/:id", createAnimalHandler)
	animalGroup.POST("/:id", updateAnimalHandler)

	shelterGroup := router.Group("/shelter/v1")
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/", getAllSheltersHandler)
	shelterGroup.POST("/:id", createShelterHandler)
	shelterGroup.POST("/:id", updateShelterHandler)
}

/* ==================================== USERS =========================================== */
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

func userSignupHandler(c *gin.Context) {
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
func getAllUsersHandler(c *gin.Context) {
}
func updateUserHandler(c *gin.Context) {
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

/* ============================== ANIMALS ======================================= */
func getAnimalByIDHandler(c *gin.Context) {
}
func getAllUsersDHandler(c *gin.Context) {
}
func createAnimalHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func updateAnimalHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}

/* ============================== SHELTERS ======================================= */
func getShelterByIDHandler(c *gin.Context) {
}
func getAllSheltersHandler(c *gin.Context) {
}
func createShelterHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func updateShelterHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
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
