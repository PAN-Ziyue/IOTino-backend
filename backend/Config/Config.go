package Config

import (
	"gorm.io/gorm"
)



var DB *gorm.DB
var DSN = "root:600019@tcp(127.0.0.1:3306)/IOTino?charset=utf8mb4&parseTime=True&loc=Local"

var PORT = 8080

var TCP_ADDR = "localhost:1883"
var KEEP_ALIVE = 20