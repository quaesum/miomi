package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"net/http"
	"strconv"
	"time"
)

func HandlerHTTP(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/login", userLoginHandler)
	api.POST("/signup", userSignupHandler)

	userGroup := api.Group("/user/v1")
	userGroup.GET("/:id", getUserByIDHandler)
	userGroup.GET("/", getAllUsersHandler)
	userGroup.POST("/signup", userSignupHandler)
	userGroup.POST("/:id", updateUserHandler)

	animalGroup := api.Group("/animal/v1")
	animalGroup.GET("/:id", getAnimalByIDHandler)
	animalGroup.GET("/", getAllAnimalsHandler)
	animalGroup.POST("/:id", createAnimalHandler)
	animalGroup.POST("/update/:id", updateAnimalHandler)

	shelterGroup := api.Group("/shelter/v1")
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/", getAllSheltersHandler)
	shelterGroup.POST("/:id", createShelterHandler)
	shelterGroup.POST("/update/:id", updateShelterHandler)

	adminGroup := api.Group("/admin/v1")
	adminGroup.Use(AdminTokenCheck()).POST("/user-update/:id", createShelterHandler)
	adminGroup.Use(AdminTokenCheck()).POST("/user-block/:id", updateShelterHandler)

	newsGroup := api.Group("/news/v1")
	newsGroup.GET("/", getAllNewsHandler)
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

/*func basicInfoHandler(c *gin.Context) {
	userinfo, err := mysql.GetUserBasicInfo(context.Background(), 1)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userinfo)
	send := fmt.Sprintf("%+v", userinfo)

	c.String(http.StatusOK, send)

}*/

func userSignupHandler(c *gin.Context) {
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	_, err := application.UserCreate(tctx, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
func userLoginHandler(c *gin.Context) {
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
	id := c.Param("id")
	uID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animals, err := application.AnimalByID(tctx, uID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
}
func getAllAnimalsHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	animals, err := application.AnimalsAll(tctx)
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
	c.JSON(200, gin.H{})
}
func getAllSheltersHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func createShelterHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func updateShelterHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}

/* =============================== NEWS ========================================= */
func getAllNewsHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	news, err := application.NewsAll(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, news)
}
