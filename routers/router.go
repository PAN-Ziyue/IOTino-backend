package routers

import (
	"IOTino/middleware/jwt"
	"IOTino/pkg/settings"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	// JWT authentication
	r.GET("/auth", GetAuth)

	// restricted operations
	group := r.Group("/api")
	group.Use(jwt.JWT())
	{
		// Device
		group.POST("/device", CreateDevice)
		group.GET("/device/:device", GetDeviceByID)

		// User
		group.POST("/register", CreateUser)
	}

	return r
}
