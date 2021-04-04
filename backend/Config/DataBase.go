package Config

import (
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN string = "root:600019@tcp(127.0.0.1:3306)/IOTino?charset=utf8mb4&parseTime=True&loc=Local"