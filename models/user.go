package models

import (
	"IOTino/pkg/e"
	"net/http"

	_ "gorm.io/driver/mysql"
)

type User struct {
	ID       uint   `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	Account  string `json:"account" gorm:"unique;size:255" validate:"required,len,account"`
	Email    string `json:"email" gorm:"unique;size:255" validate:"required,email"`
	Password string `json:"password" gorm:"size:255" swaggerignore:"true" validate:"required,len,password"`
	Verified bool   `json:"-" gorm:"default:false"`
}

type Login struct {
	Account  string `json:"account" validate:"omitempty,account"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required,password"`
	Type     string `json:"type" validate:"required"`
}

type UpdateUser struct {
	Account string `json:"account" validate:"required,len,account"`
	Email   string `json:"email" validate:"required,email"`
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

func CreateUser(user *User) e.Status {
	status := e.DefaultOk()

	err := DB.Create(&user).Error
	if err != nil {
		if CheckDuplicate(user) {
			status.Set(http.StatusBadRequest, e.DuplicateUser)
		} else {
			status.Set(http.StatusBadRequest, e.CannotCreateUser)
		}
	}

	return status
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
