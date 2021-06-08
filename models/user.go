package models

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

type User struct {
	ID       uint   `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	Account  string `json:"account" gorm:"unique;size:255"`
	Email    string `json:"email" gorm:"unique;size:255"`
	Password string `json:"password" gorm:"size:255" swaggerignore:"true"`
	Verified bool   `json:"-" gorm:"default:false"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func VerifyUser(login Login) bool {
	var users []User
	user := User{
		Email:    login.Email,
		Password: login.Password,
	}

	result := DB.Where(&user).First(&users)

	return result.RowsAffected > 0
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
