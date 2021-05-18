package models

import (
    setting "IOTino/pkg"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func init() {
    println("connect to database")

    // load database parameters
    sec, err := setting.Config.GetSection("database")

    if err != nil {
        panic(err)
    }

    DBUser := sec.Key("USER").String()
    DBPass := sec.Key("PASSWORD").String()
    DBHost := sec.Key("HOST").String()
    DBTable := sec.Key("TABLE").String()

    DSN := DBUser + ":" + DBPass + "@tcp(" + DBHost + ")/" + DBTable + "?charset=utf8mb4&parseTime=True&loc=Local"

    if DB, err = gorm.Open(mysql.Open(DSN)); err != nil {
        panic(err)
    }

    println("migrate tables")

}
