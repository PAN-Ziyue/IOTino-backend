package models

import (
	"IOTino/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Device struct {
	ID               uint    `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	UserID           uint    `json:"-" swaggerignore:"true"`
	User             User    `json:"-" gorm:"foreignKey:UserID" swaggerignore:"true"`
	Device           string  `json:"device" gorm:"unique;size:255"`
	Name             string  `json:"name" gorm:"unique;size:255"`
	Online           bool    `json:"-"`
	Alert            bool    `json:"-"`
	Count            uint    `json:"-"`
	CurrentLatitude  float64 `json:"-"`
	CurrentLongitude float64 `json:"-"`
}

// CreateDevice
// Create a device
func CreateDevice(device *Device) e.Status {
	var DuplicateDevices []Device
	status := e.New(http.StatusCreated, e.DeviceCreated)

	result := DB.Where("Device = ?", device.Device).Find(&DuplicateDevices)

	if result.RowsAffected > 0 {
		status.Set(http.StatusConflict, e.ConflictDevice)
		return status
	}

	err := DB.Create(device).Error

	if err != nil {
		status.Set(http.StatusUnprocessableEntity, e.CannotCreateDevice)
	}

	return status
}

// GetDeviceByID
// Get a device by its ID
func GetDeviceByID(DeviceID string) (e.Status, Device) {
	var device Device

	err := DB.Where("Device = ?", DeviceID).First(&device).Error
	if err != gorm.ErrRecordNotFound {
		return e.New(http.StatusOK, e.DeviceNotFound), Device{}
	}

	return e.DefaultOk(), device
}

// UpdateDevice godoc
// @Summary update a device
// @Tags Device
// @Accept  json
// @Param device path string true "device id"
// @Param name query string true "device name"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/device/{device} [PUT]
func UpdateDevice(c *gin.Context) {

}

// DeleteDevice godoc
// @Summary delete a device
// @Tags Device
// @Accept json
// @Param device path string true "device id"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/device/{device} [DELETE]
func DeleteDevice(c *gin.Context) {

}
