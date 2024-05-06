package v1

import (
	"github.com/gin-gonic/gin"
)

func HandlerHTTP(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/login", userLoginHandler)
	api.POST("/signup", userSignupHandler)

	userGroup := api.Group("/user/v1")
	userGroup.Use(UserTokenCheck())
	userGroup.GET("/info", getUserInfoHandler)
	userGroup.POST("/update", updateUserHandler)

	animalGroup := api.Group("/animal/v1")
	//animalGroup.Use(UserTokenCheck())
	animalGroup.GET("/:id", getAnimalByIDHandler)
	animalGroup.POST("/", getAnimalsHandler)
	animalGroup.POST("/add", createAnimalHandler)
	animalGroup.POST("/update/:id", updateAnimalHandler)
	animalGroup.POST("/remove/:id", removeAnimalHandler)

	shelterGroup := api.Group("/shelter/v1")
	shelterGroup.Use(UserTokenCheck())
	shelterGroup.GET("/:id", getShelterByIDHandler)
	shelterGroup.GET("/all", getAllSheltersHandler)
	shelterGroup.POST("/add", createShelterHandler)
	shelterGroup.POST("/update/:id", updateShelterHandler)

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

	serviceGroup := api.Group("/service/v1")
	serviceGroup.Use(UserTokenCheck())
	serviceGroup.GET("/:id", getServiceByIDHandler)
	serviceGroup.GET("/", getServicesHandler)
	serviceGroup.POST("/add", createServiceHandler)
	serviceGroup.POST("/update/:id", updateServiceHandler)
	serviceGroup.POST("/remove/:id", removeServiceHandler)

	productsGroup := api.Group("/products/v1")
	productsGroup.Use(UserTokenCheck())
	productsGroup.GET("/:id", getProductByIDHandler)
	productsGroup.GET("/", getProductsHandler)
	productsGroup.POST("/add", createProductHandler)
	productsGroup.POST("/update/:id", updateProductHandler)
	productsGroup.POST("/remove/:id", removeProductHandler)
}
