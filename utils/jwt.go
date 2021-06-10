package utils

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/pkg/settings"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(settings.JwtSecret)

// Claims
// JWT claims
type Claims struct {
	ID    uint
	Email string
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := e.DefaultOk()

		token, err := c.Cookie("token")

		if token == "" {
			status.Set(http.StatusUnauthorized, e.Unauthorized)
			c.JSON(status.Code, gin.H{"msg": status.Msg})
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			status.Set(http.StatusUnauthorized, e.ParseTokenError)
		} else if time.Now().Unix() > claims.ExpiresAt {
			status.Set(http.StatusUnauthorized, e.AuthTimeout)
		}

		if status.Code != http.StatusOK {
			c.JSON(status.Code, gin.H{"msg": status.Msg})
			c.Abort()
			return
		}

		user, err := models.GetUserByID(claims.ID)
		if err != nil {
			status.Set(http.StatusNoContent, e.UserNotFound)
			c.JSON(status.Code, gin.H{"msg": status.Msg})
			c.Abort()
			return
		}

		c.Set("auth", user)
		c.Next()
	}
}

func GenerateToken(id uint, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
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
