package controllers

import (
	model "franchise-web/app/models"
	data "franchise-web/config/initializers"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetAllCustomers godoc
// @Summary      Test
// @Description  test
// @Tags         Test
// @Accept       json
// @Produce      json
// @Router       /api/v1/ping [get]
func TestController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request data",
		})
		return
	}

	user.Password = user.Password + "safety"
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encrypt password",
		})
		return
	}

	log.Println(string(hashedPassword))
	user.Password = string(hashedPassword)

	if err := data.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "User registered successfully",
		"password": string(hashedPassword),
	})
}
