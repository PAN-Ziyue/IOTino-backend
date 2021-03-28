package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	host     string
	port     int
	db_name  string
	user     string
	password string
}

func BuildConfig() *DBConfig {
	config := DBConfig{
		host:     "localhost",
		port:     3306,
		db_name:  "IOTino",
		user:     "root",
		password: "600019",
	}
	return &config
}

func DataBaseURL(db_config *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db_config.user,
		db_config.password,
		db_config.host,
		db_config.port,
		db_config.db_name,
	)
}
