package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"task_system_go/database"
	"task_system_go/models"
	"task_system_go/token"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if hashedPassword, err := models.HashPassword(user.Password); err == nil {
		user.Password = hashedPassword
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't hash password"})
		c.Abort()
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, gin.H{"message": "User was created"})
}

func Login(c *gin.Context) {
	var user models.User
	var loginData LoginPayload
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	res := database.Database.Where("Username = ?", loginData.Username).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
		c.Abort()
		return
	}

	if hashedPassword, err := models.HashPassword(loginData.Password); err != nil || !user.ValidatePassword(hashedPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		c.Abort()
		return
	}

	signedToken, refreshToken, err := token.GenerateAllTokens(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tokenResponse := TokenResponse{
		SignedToken:  signedToken,
		RefreshToken: refreshToken}
	c.JSON(http.StatusOK, tokenResponse)

}
