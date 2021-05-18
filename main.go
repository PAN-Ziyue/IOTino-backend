package main

import (
    _ "IOTino/docs" // docs is generated by Swag CLI, you have to import it.
    "IOTino/models"

    setting "IOTino/pkg"
    "IOTino/pkg/MQTT"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

// @title IOTino
// @version 0.0.1
// @description The worst IOT website in Laohe Mountain Institute of Technology
// @host localhost:8080
func main() {

    var err error

    setting.Init()

    println("set up MQTT broker")
    go MQTT.MQTT()

    println(">>> Stage 3: setup web app server")
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    group := r.Group("/api")
    {
        group.POST("/register", models.CreateUser)
    }

    if err := r.Run(); err != nil {
        fmt.Println(err)
        return
    }
}