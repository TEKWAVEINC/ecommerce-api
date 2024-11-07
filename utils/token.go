package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))


type Claims struct {
	UserID  uint
	Role	string
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	expirationTime := time.Now().Add(24 *time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:	role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}