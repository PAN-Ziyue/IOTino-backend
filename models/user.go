package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

type User struct {
	ID       uint   `gorm:"primaryKey" swaggerignore:"true"`
	Account  string `json:"account" gorm:"unique;size:255"`
	Email    string `json:"email" gorm:"unique;size:255"`
	Password string `json:"password" gorm:"size:255" swaggerignore:"true"`
	Verified bool   `gorm:"default:false"`
}

// TODO add error handling

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
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	print("ok")
	c.String(http.StatusOK, "ok")
}

func VerifyUser(c *gin.Context) {

}

// UpdatePassword godoc
// @Summary update a user's password
// @Tags User
// @Accept  json
// @Param password query string true "password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [PUT]
func UpdatePassword(c *gin.Context) {

}

// DeleteUser godoc
// @Summary delete a user
// @Tags User
// @Accept  json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [DELETE]
func DeleteUser(c *gin.Context) {

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
