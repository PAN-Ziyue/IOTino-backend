package main

import (
	"IOTino/Config"
	"IOTino/Models"
	"IOTino/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()

	r.Run()
}
