package authHelper

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	ID    uint
	Name  string
	Email string
}

func (user *User) GenerateToken() (string, error) {
	jwtSecret := os.Getenv("SECRET_KEY")
	if jwtSecret == "" {
		return "", errors.New("JWT secret key not found")
	}

	now := time.Now()
	jwtId := uuid.New().String()
	claims := Claims{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Audience:  "franchise-web-user",
			ExpiresAt: now.Add(time.Hour * 24 * 7).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "franchise-web-2024",
			NotBefore: now.Add(5 * time.Second).Unix(),
			Subject:   user.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("token: ", token)
	return token.SignedString([]byte(jwtSecret))
}
