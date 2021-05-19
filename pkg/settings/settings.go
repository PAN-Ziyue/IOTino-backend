package settings

import (
	"github.com/go-ini/ini"

	"time"
)

// .ini file
var Config *ini.File

// server
var RunMode string
var HttpPort string
var MQTTPort int
var ReadTimeOut time.Duration
var WriteTimeOut time.Duration
var JwtSecret string

// MQTT
var KEEP_ALIVE int

func InitSettings() {
	var err error

	println("load init file")

	// load .ini file
	Config, err = ini.Load("config/IOTino.ini")
	if err != nil {
		panic(err)
	}

	sec, err := Config.GetSection("server")
	if err != nil {
		panic(err)
	}

	HttpPort = ":" + sec.Key("HTTP_PORT").String()
	RunMode = sec.Key("RUN_MODE").String()

}

var TCP_ADDR = "localhost:1883"

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
