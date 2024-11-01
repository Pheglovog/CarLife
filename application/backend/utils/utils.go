package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

func GenerateJWT(userID string, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"userType": userType,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}

func CheckPassword(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func ParseJWT(tokenString string) (string, string, error) {
	// Remove Bearer from tokenString
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", "", err
	}
	// Get userid and userType from token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(string)
		if !ok {
			return "", "", errors.New("userID not found in token")
		}
		userType, ok := claims["userType"].(string)
		if !ok {
			return "", "", errors.New("userType not found in token")
		}
		return userID, userType, nil
	}

	return "", "", errors.New("token invalid")
}
