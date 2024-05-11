package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"madmax/internal/application"
	"madmax/internal/entity"
	"madmax/internal/utils"
	"net/http"
	"strconv"
	"time"
)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
	return
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
	return
}
func userLoginHandler(c *gin.Context) {
	var ul entity.UserLogInRequest
	if err := c.ShouldBindJSON(&ul); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Проверьте данные"})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*5)
	token, err := application.LogIn(tctx, ul.Email, ul.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	c.JSON(200, gin.H{"data": token})
	return
}

func getAllUsersHandler(c *gin.Context) {
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Second*2)
	users, err := application.GetAllUsers(tctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"data": users})
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
	return
}

func getUserInfoHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Background()
	tctx, _ := context.WithTimeout(ctx, time.Minute*5)

	user, err := application.UserByID(tctx, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
	return
}

func verifyEmailSendHandler(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = application.VerifyUserEmail(context.Background(), userID)
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, gin.H{
		"data": "ok",
	})
}

func verifyEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.Redirect(http.StatusFound, "https://app.dealingi.by/email-confirm-restricted/")
	}
	err := application.VerifyEmail(context.Background(), token)
	if err != nil {
		c.Redirect(http.StatusFound, "https://app.dealingi.by/email-confirm-restricted/")
	}
	c.Redirect(http.StatusFound, "https://app.dealingi.by/email-confirm-success")
}
