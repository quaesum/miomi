package v1

import (
	"github.com/gin-gonic/gin"
)

func HandlerHTTP(router *gin.Engine) {

	products := NewProductsHttp()
	animals := NewAnimalsHttp()
	services := NewServicesHttp()

	api := router.Group("/api")
	api.POST("/login", userLoginHandler)
	api.POST("/signup", userSignupHandler)

	userGroup := api.Group("/user/v1")
	//userGroup.Use(UserTokenCheck())
	userGroup.GET("/info", getUserInfoHandler)
	userGroup.POST("/update", updateUserHandler)
	userGroup.GET("/info/:id", getUserByIDHandler)
	userGroup.POST("/verify_email", verifyEmailSendHandler)
	userGroup.GET("/verify_email", verifyEmailHandler)

	animalGroup := api.Group("/animal/v1")
	//animalGroup.Use(UserTokenCheck())
	animalGroup.GET("/:id", animals.GetByID)
	animalGroup.POST("/", animals.GetAll)
	animalGroup.POST("/add", animals.Create)
	animalGroup.POST("/update/:id", animals.Update)
	animalGroup.POST("/remove/:id", animals.Remove)

	shelterGroup := api.Group("/shelter/v1")
	shelterGroup.Use(UserTokenCheck())
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/all", getAllSheltersHandler)
	shelterGroup.POST("/add", createShelterHandler)
	shelterGroup.POST("/update/:id", updateShelterHandler)
	shelterGroup.POST("/remove/:id", removeShelterHandler)

	adminGroup := api.Group("/admin/v1")
	adminGroup.Use(AdminTokenCheck())
	adminGroup.POST("/user-update/:id", createShelterHandler)
	adminGroup.POST("/user-block/:id", updateShelterHandler)
	adminGroup.GET("/allUsers", getAllUsersHandler)
	adminGroup.POST("/:id", getUserByIDHandler)

	newsGroup := api.Group("/news/v1")
	//newsGroup.Use(UserTokenCheck())
	newsGroup.GET("/", getAllNewsHandler)
	newsGroup.POST("/add", createNewsHandler)
	newsGroup.POST("/update/:id", updateNewsHandler)
	newsGroup.POST("/remove/:id", removeNewsHandler)

	fileGroup := api.Group("/file/v1")
	//fileGroup.Use(UserTokenCheck())
	fileGroup.POST("/add", attachAnimalFileHandler)
	fileGroup.POST("/addNews", attachNewsFileHandler)
	fileGroup.POST("/addService", attachServiceFileHandler)
	fileGroup.POST("/addProduct", attachProductFileHandler)
	fileGroup.GET("/getUrl", getAllFileNamesAndIdsHandler)

	serviceGroup := api.Group("/services/v1")
	//serviceGroup.Use(UserTokenCheck())
	serviceGroup.GET("/:id", services.GetByID)
	serviceGroup.POST("/", services.GetAll)
	serviceGroup.POST("/add", services.Create)
	serviceGroup.POST("/update/:id", services.Update)
	serviceGroup.POST("/remove/:id", services.Remove)

	productsGroup := api.Group("/products/v1")
	//productsGroup.Use(UserTokenCheck())
	productsGroup.GET("/:id", products.GetByID)
	productsGroup.POST("/", products.GetAll)
	productsGroup.POST("/add", products.Create)
	productsGroup.POST("/update/:id", products.Update)
	productsGroup.POST("/remove/:id", products.Remove)
}

type Controller interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Remove(c *gin.Context)
}
