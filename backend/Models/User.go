package Models

import (
	"IOTino/Config"
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "gorm.io/driver/mysql"
	"net/http"
)

type User struct {
	ID       uint         `gorm:"primaryKey"`
	Account  string       `json:"account" gorm:"unique"`
	Email    string       `json:"email" gorm:"unique"`
	Password string       `json:"password"`
	Verified sql.NullBool `gorm:"default:false"`
}

// TODO add error handling
func CreateUser(c echo.Context) error {
	var user User
	_ = c.Bind(&user)

	if err := Config.DB.Create(&user).Error; err != nil {
		print("error found")
		return c.String(http.StatusUnauthorized, "error")
	}

	print("ok")
	return c.String(http.StatusOK, "ok")
}


func ValidateUser(c echo.Context) error {
	// TODO
}