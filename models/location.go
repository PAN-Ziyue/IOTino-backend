package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Location struct {
	ID        uint   `gorm:"primaryKey" swaggerignore:"true"`
	DeviceID  uint   `swaggerignore:"true"`
	Device    Device `gorm:"foreignKey:DeviceID" swaggerignore:"true"`
	Latitude  float64
	Longitude float64
	Time      time.Time
}

// GetAllLocationSequence godoc
// @Summary get all device's history location sequence
// @Tags Location
// @Accept  json
// @Success 200 {array} Location
// @Failure 400 {string} string "error"
// @Router /api/locations [GET]
func GetAllLocationSequence(c *gin.Context) {

}

// GetDeviceLocationSequence godoc
// @Summary get certain device's history location sequence
// @Tags Location
// @Accept  json
// @Param device path string true "device id"
// @Success 200 {array} Location
// @Failure 400 {string} string "error"
// @Router /api/location/{device} [GET]
func GetDeviceLocationSequence(c *gin.Context) {

}
