package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
)

type Auth struct {
	Account string `validate:"required,gt=10"`
	Email   string `validate:"required,email"`
}

var validate = validator.New()

func Login(c *gin.Context) {
	account := c.Query("account")
	email := c.Query("email")
	status := e.DefaultOk()

	// create jwt auth
	auth := Auth{Account: account, Email: email}

	// authenticate
	err := validate.Struct(auth)
	if err != nil {
		status = e.New(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// make data
	data := make(map[string]interface{})

	// check exist
	isExist := models.CheckAuth(account, email)

	// TODO handle auth error
	if isExist {
		token, err := utils.GenerateToken(account, email)
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
		c.JSON(status.Code, gin.H{"error": status.Msg})
		return
	}

	// create user
	if err := models.DB.Create(&user).Error; err != nil {
		status := e.DefaultError()

		var mysqlErr mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			status.Msg = e.DuplicateUser
		}

		c.JSON(status.Code, gin.H{"error": status.Msg})
		return
	}

	c.String(http.StatusOK, "ok")
}
