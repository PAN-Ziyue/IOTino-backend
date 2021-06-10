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

func VerifyUser(login Login) User {
	var user User

	DB.Where(&User{Email: login.Email, Password: login.Password}).First(&user)

	return user
}

func CheckDuplicate(user *User) bool {
	var emailUser User
	var accountUser User

	emailResult := DB.Where(&User{Email: user.Email}).First(&emailUser)
	accountResult := DB.Where(&User{Account: user.Account}).First(&accountUser)

	return emailResult.RowsAffected > 0 || accountResult.RowsAffected > 0
}

func GetUserByEmail(email string) (User, error) {
	var user User

	// query
	err := DB.Where(&User{Email: email}).First(&user).Error

	return user, err
}

func GetUserByID(id uint) (User, error) {
	var user User

	err := DB.First(&user, id).Error

	return user, err
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
