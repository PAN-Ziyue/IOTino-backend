package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Auth struct {
	Account string `validate:"required,gt=10"`
	Email   string `validate:"required,email"`
}

var validate = validator.New()

func GetAuth(c *gin.Context) {
	account := c.Query("account")
	email := c.Query("email")
	status := e.DefaultStatus()

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
