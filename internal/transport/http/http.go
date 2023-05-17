package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
	"net/http"
	"strconv"
	"time"
)

func HTTPHandler(router *gin.Engine) {
	api := router.Group("/api")
	userGroup := api.Group("/user/v1")
	userGroup.GET("/:id", getUserByIDHandler)
	userGroup.GET("/", getAllUsersHandler)
	userGroup.Use(AdminTokenCheck()).POST("/signup", userSignupHandler)
	userGroup.Use(AdminTokenCheck()).POST("/:id", updateUserHandler)

	animalGroup := api.Group("/animal/v1")
	animalGroup.GET("/:id", getAnimalByIDHandler)
	animalGroup.GET("/", getAllUsersDHandler)
	animalGroup.Use(AdminTokenCheck()).POST("/:id", createAnimalHandler)
	animalGroup.Use(AdminTokenCheck()).POST("/update/:id", updateAnimalHandler)

	shelterGroup := api.Group("/shelter/v1")
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/", getAllSheltersHandler)
	shelterGroup.Use(AdminTokenCheck()).POST("/:id", createShelterHandler)
	shelterGroup.Use(AdminTokenCheck()).POST("/update/:id", updateShelterHandler)
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
	application.UserByID(tctx, uID)
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
	animals, err := application.AnimalsAll(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
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
