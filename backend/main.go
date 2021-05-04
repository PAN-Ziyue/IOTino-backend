package main

import (
	"IOTino/Config"
	"IOTino/Models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

		group.POST("/register", Models.CreateUser)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
