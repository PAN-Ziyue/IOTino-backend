package Models

import "github.com/gin-gonic/gin"

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

// GetDeviceByID godoc
// @Summary get a device by ID
// @Tags Device
// @Accept  json
// @Param device path string true "device id"
// @Success 200 {object} Device
// @Failure 400 {string} string "error"
// @Router /api/device/{device} [GET]
func GetDeviceByID(c *gin.Context) {

}

// GetDevices godoc
// @Summary get all devices
// @Tags Device
// @Accept  json
// @Success 200 {array} Device
// @Failure 400 {string} string "error"
// @Router /api/devices [GET]
func GetDevices(c *gin.Context) {

}

// CreateDevice godoc
// @Summary create a device
// @Tags Device
// @Accept  json
// @Param device query string true "device id"
// @Param name query string true "device name"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/device [POST]
func CreateDevice(c *gin.Context) {

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
