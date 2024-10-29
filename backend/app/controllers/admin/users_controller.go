package controllers

import (
	authHelper "franchise-web/app/helper/admin"
	model "franchise-web/app/models"
	data "franchise-web/config/initializers"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Test godoc
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

// Register godoc
// @Summary      Register a new User
// @Description  Register a new User with the provided credentials
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        User  body  services.RegisterRequest  true  "User information"
// @Success 200  {object} services.SuccessfulSignup "User registered successfully"
// @Failure 400  {object} services.BadRequest "Bad request response"
// @Failure 500  {object} services.ServerError "Internal server error response"
// @Router       /api/v1/register [post]
func Register(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request data",
		})
		return
	}

	// set password
	user.Password = user.Password + "safety"
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encrypt password",
		})
		return
	}

	user.Password = string(hashedPassword)

	if err := data.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// Login godoc
// @Summary      Login
// @Description  Authenticate a user and generate a JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        User  body  services.LoginRequest  true  "User credentials"
// @Success 200 {object} services.SuccessfulLoginResponse "Successful login response"
// @Failure 401 {object} services.HttpError "Incorrect email or password"
// @Failure 500 {object} services.ServerError "Error generating token"
// @Router       /api/v1/login [post]
func Login(c *gin.Context) {
	var user model.User
	var foundUser model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request data",
		})
		return
	}

	if err := data.DB.Where("email=?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect email",
		})
		return
	}

	user.Password = user.Password + "safety"
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect Password",
		})
		return
	}
	authUser := authHelper.User{
		ID:    foundUser.ID,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}

	token, err := authUser.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User login successfully",
		"token":   token,
	})
}

// Logout godoc
// @Summary      Logout a User
// @Description  Clear the authentication token and log out the User
// @Tags         User
// @Accept       json
// @Produce      json
// @Success 200  {object} services.SuccessfulLogoutResponse "Logged out successfully"
// @Router       /api/v1/logout [delete]
func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not provided",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signed out successfully",
	})
}
