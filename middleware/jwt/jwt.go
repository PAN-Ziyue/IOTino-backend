package jwt

import (
	"IOTino/pkg/e"
	"IOTino/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		status := e.DefaultStatus()
		token := c.Query("token")

		if token == "" {
			status.Set(http.StatusBadRequest, e.BadParameter)
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				status.Set(http.StatusUnauthorized, e.ParseTokenError)
			} else if time.Now().Unix() > claims.ExpiresAt {
				status.Set(http.StatusUnauthorized, e.AuthTimeout)
			}
		}

		if status.Code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  status.Msg,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
