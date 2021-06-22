package routers

import (
	"IOTino/pkg/settings"
	"IOTino/utils"

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

	// user
	r.POST("/api/login", Login)
	r.POST("/api/register", CreateUser)

	// restricted operations
	group := r.Group("/api")
	group.Use(utils.JWT())
	{
		// Device
		group.GET("/currentUser", CurrentUser)
		group.GET("/dashboard", GetDashboard)
		group.PUT("/users", UpdateUser)
		group.GET("/logout", LogoutUser)

		group.POST("/devices", CreateDevice)
		group.GET("/devices", GetDevices)
		group.DELETE("/devices", DeleteDevice)
		group.PUT("/devices", UpdateDevice)
	}

	return r
}
