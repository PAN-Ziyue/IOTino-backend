package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

type DBConfig struct {
	host     string
	port     int
	db_name  string
	user     string
	password string
}

func build_config() *DBConfig {
	Config := DBConfig{
		host:     "localhost",
		port:     27017,
		db_name:  "first_go",
		user:     "admin",
		password: "1234",
	}
	return &Config
}

func database_url(db_config *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db_config.user,
		db_config.password,
		db_config.host,
		db_config.port,
		db_config.db_name,
	)
}
