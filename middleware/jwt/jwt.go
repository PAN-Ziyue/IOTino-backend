package jwt

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := e.DefaultOk()

		token, err := c.Cookie("token")

		if token == "" {
			status.Set(http.StatusBadRequest, e.BadParameter)
			c.JSON(status.Code, gin.H{"msg": status.Msg})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
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
