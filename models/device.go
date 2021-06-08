package models

import (
	"IOTino/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Device struct {
	ID               uint   `gorm:"primaryKey" swaggerignore:"true"`
	UserID           uint   `swaggerignore:"true"`
	User             User   `gorm:"foreignKey:UserID" swaggerignore:"true"`
	Device           string `json:"device" gorm:"unique;size:255"`
	Name             string `json:"name" gorm:"unique;size:255"`
	Online           bool
	Alert            bool
	Count            uint
	CurrentLatitude  float64
	CurrentLongitude float64
}

// CreateDevice
// Create a device
func CreateDevice(device *Device) e.Status {
	var DuplicateDevices []Device
	err := DB.Where("Device = ?", device.Device).Find(&DuplicateDevices).Error
	if err != gorm.ErrRecordNotFound {
		return e.New(http.StatusConflict, e.ConflictDevice)
	}

	DB.Create(Device{})
	return e.New(http.StatusCreated, e.DeviceCreated)
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
