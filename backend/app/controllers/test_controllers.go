package controllers

import "github.com/gin-gonic/gin"

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
