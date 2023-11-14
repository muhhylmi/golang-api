package utils

import (
	"golang-api/config"
	"golang-api/modules/users/models/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user *domain.Users) (string, error) {
	config := config.GetConfig()
	// Membuat token
	token := jwt.New(jwt.SigningMethodHS256)
	var secretKey = []byte(config.JWT_SECRET_KEY)

	// Menambahkan klaim ke token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["gender"] = user.Gender
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Menandatangani token dengan kunci rahasia
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
