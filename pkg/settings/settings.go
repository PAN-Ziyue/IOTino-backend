package settings

import (
	"github.com/go-ini/ini"

	"time"
)

// Config
// load from IOTino.ini
var Config *ini.File

// mqtt
var MQTTAddr string
var MQTTPort string
var MQTTHost string
var KeepAlive int

// server
var RunMode string
var HTTPPort string
var ReadTimeOut time.Duration
var WriteTimeOut time.Duration
var JwtSecret string

func InitSettings() {
	var err error

	println("load init file")

	// load .ini file
	Config, err = ini.Load("config/IOTino.ini")
	if err != nil {
		panic(err)
	}

	MQTTSection, err := Config.GetSection("MQTT")
	if err != nil {
		panic(err)
	}

	MQTTPort = MQTTSection.Key("MQTTPort").String()
	MQTTAddr = MQTTSection.Key("MQTTAddr").String()
	MQTTHost = MQTTAddr + ":" + MQTTPort
	KeepAlive, err = MQTTSection.Key("KeepAlive").Int()

	if err != nil {
		panic(err)
	}

	ServerSection, err := Config.GetSection("Server")
	if err != nil {
		panic(err)
	}

	HTTPPort = ":" + ServerSection.Key("HTTPPort").String()
	RunMode = ServerSection.Key("RunMode").String()
}
