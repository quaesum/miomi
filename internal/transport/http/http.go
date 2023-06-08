package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"madmax/internal/application"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func HandlerHTTP(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/login", userLoginHandler)
	api.POST("/signup", userSignupHandler)

	userGroup := api.Group("/user/v1")
	userGroup.GET("/:id", getUserByIDHandler)
	userGroup.GET("/all", getAllUsersHandler)
	userGroup.GET("/info", getUserInfoHandler)
	userGroup.POST("/update", updateUserHandler)

	animalGroup := api.Group("/animal/v1")
	animalGroup.GET("/:id", getAnimalByIDHandler)
	animalGroup.GET("/", getAllAnimalsHandler)
	animalGroup.POST("/add", createAnimalHandler)
	animalGroup.POST("/update/:id", updateAnimalHandler)
	animalGroup.POST("/remove/:id", removeAnimalHandler)

	shelterGroup := api.Group("/shelter/v1")
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/all", getAllSheltersHandler)
	shelterGroup.POST("/add", createShelterHandler)
	shelterGroup.POST("/update/:id", updateShelterHandler)

	adminGroup := api.Group("/admin/v1")
	adminGroup.Use(AdminTokenCheck()).POST("/user-update/:id", createShelterHandler)
	adminGroup.Use(AdminTokenCheck()).POST("/user-block/:id", updateShelterHandler)

	newsGroup := api.Group("/news/v1")
	newsGroup.GET("/", getAllNewsHandler)
	newsGroup.POST("/add", createNewsHandler)
	newsGroup.POST("/remove/:id", removeNewsHandler)

	fileGroup := api.Group("/file/v1")
	fileGroup.POST("/add", attachFileHandler)
	fileGroup.POST("/addNews", attachNewsFileHandler)
	fileGroup.GET("/getUrl", getAllFileNamesAndIdsHandler)
}

/* ==================================== USERS =========================================== */
func getUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*2)
	user, err := application.UserByID(tctx, userID)
	if err != nil {
		id := c.Param("id")
		uID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := context.Background()
		tctx, _ := context.WithTimeout(ctx, time.Minute*2)
		application.UserByID(tctx, uID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func userSignupHandler(c *gin.Context) {
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	token, err := application.UserCreate(tctx, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": token})

}
func userLoginHandler(c *gin.Context) {
	var ul entity.UserLogInRequest
	if err := c.ShouldBindJSON(&ul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	token, err := application.LogIn(tctx, ul.Email, ul.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": token})
}

func getAllUsersHandler(c *gin.Context) {
}
func updateUserHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ucr entity.UserCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.UserUpdate(tctx, userID, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func getUserInfoHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)

	user, err := application.UserByID(tctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})

}

/* ============================== ANIMALS ======================================= */
func getAnimalByIDHandler(c *gin.Context) {
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animals, err := application.AnimalByID(tctx, animalID)
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
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ucr entity.AnimalCreateRequest
	if err := c.ShouldBindJSON(&ucr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	animalID, err := application.AnimalCreate(tctx, userID, &ucr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": animalID})
}
func updateAnimalHandler(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func removeAnimalHandler(c *gin.Context) {
	_, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	animalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveAnimalByID(tctx, animalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

/* ============================== SHELTERS ======================================= */
func getShelterByIDHandler(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*2)
	user, err := application.ShelterByID(tctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}
func getAllSheltersHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*15)
	animals, err := application.SheltersAll(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, animals)
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

func createNewsHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var ncr entity.NewsCreateRequest
	if err := c.ShouldBindJSON(&ncr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	newsId, err := application.NewsCreate(tctx, userID, &ncr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": newsId})
}

func removeNewsHandler(c *gin.Context) {
	_, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	newsId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	err = application.RemoveNewsById(tctx, newsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

/* =============================== FILES ========================================= */
func attachFileHandler(c *gin.Context) {
	_, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	newFileName := uuid.New().String()
	ctx := context.Background()
	f, err := file.Open()

	resp, err := application.AddAnimalsFile(ctx, file.Size, fmt.Sprintf("%s%s", newFileName, strings.ToLower(filepath.Ext(file.Filename))), f)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error add file": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
}

func attachNewsFileHandler(c *gin.Context) {
	_, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	newFileName := uuid.New().String()
	ctx := context.Background()
	f, err := file.Open()

	resp, err := application.AddNewsFile(ctx, file.Size, fmt.Sprintf("%s%s", newFileName, strings.ToLower(filepath.Ext(file.Filename))), f)

	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error add file": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"data": resp,
	})
}

func getAllFileNamesAndIdsHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*10)
	photos, err := application.GetFileNameAndId(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, photos)
}
