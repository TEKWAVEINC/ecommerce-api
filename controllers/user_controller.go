package controllers

import (
	"log"
	"net/http"
	"ecommerce-api/config"
	"ecommerce-api/models"
	"ecommerce-api/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


func RegisterUser(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	user := models.User{Email: input.Email}
	if err := user.HashPassword(input.Password); err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error: password hashing failed"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Printf("Error registering user: %v", err)
		c.JSON(http.StatusConflict, gin.H{"error": "User registration failed, user might already exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}


func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: user not found"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		log.Printf("Password check failed: %v", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Authentication failed: incorrect password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error: token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
