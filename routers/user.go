package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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

	// make data
	data := make(map[string]interface{})

	// check exist
	exist := models.VerifyUser(login)

	if exist {
		token, err := utils.GenerateToken(login.Email)
		if err != nil {
			status.Set(http.StatusUnauthorized, e.CannotGenToken)
		} else {
			data["token"] = token
		}
	} else {
		status.Set(http.StatusUnauthorized, e.WrongAccount)
	}

	c.JSON(status.Code, gin.H{
		"msg":  status.Msg,
		"data": data,
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

	// bind parameter
	if err := c.ShouldBindJSON(&user); err != nil {
		status := e.New(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// create user
	if err := models.DB.Create(&user).Error; err != nil {
		status := e.DefaultError()

		var mysqlErr mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			status.Msg = e.DuplicateUser
		}

		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.String(http.StatusOK, "ok")
}
