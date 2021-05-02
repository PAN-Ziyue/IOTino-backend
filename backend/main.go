package main

import (
	"IOTino/Config"
	"IOTino/Models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var err error

func main() {

	Config.DB, err = gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})

	Config.DB.AutoMigrate(&Models.User{})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	group := e.Group("/api")
	{

		group.POST("/register", CreateUser)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func newHTTPError(code int, key string, msg string) *httpError {
	return &httpError{
		code:    code,
		Key:     key,
		Message: msg,
	}
}

func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

//CreateUser ... Create User
func CreateUser(c echo.Context) error {
	var user Models.User
	_ = c.Bind(&user)

	if err = Config.DB.Create(&user).Error; err != nil {
		print("error found")
		return c.String(http.StatusUnauthorized, "error")
	}

	print("ok")
	return c.String(http.StatusOK, "ok")
}
