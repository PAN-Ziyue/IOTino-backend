package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/pkg/settings"
	"IOTino/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login
	status := e.DefaultOk()

	if err := c.ShouldBindJSON(&login); err != nil {
		println("[LOG] invalid parameter")
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// check exist
	user := models.VerifyUser(login)

	if (models.User{}) == user {
		status.Set(http.StatusUnauthorized, e.WrongAccount)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// make data
	token, err := utils.GenerateToken(user.ID, login.Email)

	if err != nil {
		status.Set(http.StatusUnauthorized, e.CannotGenToken)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.SetCookie("token",
		token,
		3600,
		"/api/",
		settings.Domain,
		false,
		true,
	)

	c.JSON(status.Code, gin.H{
		"msg":   status.Msg,
		"token": token,
	})
}

// CreateUser godoc
// @Summary create a user
// @Tags User
// @Accept  json
// @Param account query string true "account"
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [POST]
func CreateUser(c *gin.Context) {
	var user models.User
	status := e.DefaultOk()

	// bind parameter
	if err := c.ShouldBindJSON(&user); err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// create user
	if err := models.DB.Create(&user).Error; err != nil {
		if models.CheckDuplicate(&user) {
			status.Set(http.StatusBadRequest, e.DuplicateUser)
		}

		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		status.Set(http.StatusUnauthorized, e.CannotGenToken)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.SetCookie("token",
		token,
		3600,
		"/api/",
		settings.Domain,
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"token": token,
	})
}

// DeleteUser godoc
// @Summary delete a user
// @Tags User
// @Accept  json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [DELETE]
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		status := e.New(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.String(http.StatusOK, "ok")
}

// GetUser godoc
// @Summary get a user's specification
// @Tags User
// @Accept  json
// @Success 200 {object} User
// @Failure 400 {string} string "error"
// @Router /api/user [GET]
func GetUser(c *gin.Context) {

}
