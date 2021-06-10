package main

import (
	_ "IOTino/docs" // docs is generated by Swag CLI, you have to import it.
	"IOTino/models"
	"IOTino/pkg/mqtt"
	"IOTino/pkg/settings"
	"IOTino/routers"
	"IOTino/utils"
	"fmt"
)

// @title IOTino
// @version 0.0.1
// @description The worst IOT website in Laohe Mountain Institute of Technology
// @host localhost:8080
func main() {

	var err error

	settings.InitSettings()
	models.InitDB()
	utils.InitValidator()

	println("set up mqtt broker")
	go mqtt.MQTT()

	println("setup web app server")
	r := routers.InitRouter()

	if err = r.Run(settings.HTTPPort); err != nil {
		fmt.Println(err)
		return
	}
}
