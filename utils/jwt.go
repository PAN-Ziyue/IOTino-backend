package utils

import (
	"IOTino/pkg/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(settings.JwtSecret)

// Claims
// JWT claims
type Claims struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(account string, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		account,
		email,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "IOTino",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
