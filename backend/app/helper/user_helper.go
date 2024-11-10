package helper

import (
	"github.com/gin-gonic/gin"

	model "franchise-web/app/models"
	"franchise-web/config/initializers"
)

func FindUserByID(c *gin.Context) (*model.User, error) {
	var user model.User

	id := c.Param("id")

	if err := initializers.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
