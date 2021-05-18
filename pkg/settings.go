package setting

import (
    "github.com/go-ini/ini"

    "time"
)

// .ini file
var Config *ini.File

// server
var HTTPPort int
var ReadTimeOut time.Duration
var WriteTimeOut time.Duration
var JwtSecret string

func Init() {
    var err error

    println("load init file")

    // load .ini file
    if Config, err = ini.Load("config/IOTino.ini"); err != nil {
        panic(err)
    }

    if HTTPPort, err = Config.Section("server").Key("HTTPPort").Int(); err != nil {
        panic(err)
    }
}

var TCP_ADDR = "localhost:1883"
var KEEP_ALIVE = 20

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
