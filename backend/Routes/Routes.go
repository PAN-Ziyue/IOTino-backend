package Routes

import (
	"IOTino/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.POST("register", Controllers.CreateUser)
	}
	return r
}
