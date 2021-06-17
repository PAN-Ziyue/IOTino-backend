package models

import (
	"IOTino/pkg/settings"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	println("connect to database")

	// load database parameters
	sec, err := settings.Config.GetSection("Database")

	if err != nil {
		panic(err)
	}

	DBUser := sec.Key("USER").String()
	DBPass := sec.Key("PASSWORD").String()
	DBHost := sec.Key("HOST").String()
	DBTable := sec.Key("TABLE").String()

	DSN := DBUser + ":" + DBPass + "@tcp(" + DBHost + ")/" + DBTable + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(DSN))
	if err != nil {
		log.Panicln("Cannot open database due to:", err)
	}

	println("migrate tables")

	err = DB.AutoMigrate(User{}, Device{}, Location{})
	if err != nil {
		log.Panicln("Cannot migrate tables due to:", err)
	}

	DB.Model(&Device{}).Update("status", offline)
}
