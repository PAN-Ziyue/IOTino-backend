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
	r.GET("/login", Login)
	r.POST("/register", CreateUser)

	// restricted operations
	group := r.Group("/api")
	group.Use(utils.JWT())
	{
		// Device
		group.POST("/device", CreateDevice)
		group.GET("/devices", GetDevices)
		group.GET("/device/:device", GetDeviceByID)
	}

	return r
}

//func InitRouter() *gin.Engine {
//	r := gin.New()
//
//	r.Use(gin.Logger())
//	r.Use(gin.Recovery())
//	r.Use(TlsHandler())
//
//	gin.SetMode(settings.RunMode)
//
//	r.GET("/swagger/*any",
//		ginSwagger.WrapHandler(swaggerFiles.Handler))
//
//	// user
//	r.GET("/login", Login)
//	r.POST("/register", CreateUser)
//
//
//	// restricted operations
//	group := r.Group("/api")
//	group.Use(jwt.JWT())
//	{
//		// Device
//		group.POST("/device", CreateDevice)
//		group.GET("/device/:device", GetDeviceByID)
//	}
//
//	return r
//}
//
//
//func TlsHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		secureMiddleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     settings.SSLHost,
//		})
//		err := secureMiddleware.Process(c.Writer, c.Request)
//
//		// If there was an error, do not continue.
//		if err != nil {
//			return
//		}
//
//		c.Next()
//	}
//}
