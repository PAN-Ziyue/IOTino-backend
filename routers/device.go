package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var deviceJSON models.DeviceJSON
	var status = e.DefaultOk()

	// bind model
	err := c.BindJSON(&deviceJSON)
	if err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	device := models.Device{
		Device: deviceJSON.Device,
		Name:   deviceJSON.Name,
	}
	// get user metadata

	auth_user, exist := c.Get("auth")

	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := auth_user.(models.User)

	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	device.UserID = user.ID
	device.User = user

	status = models.CreateDevice(&device)

	c.JSON(status.Code, gin.H{"msg": status.Msg})
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
	var data models.Device
	var status = e.DefaultOk()
	var DeviceID string = c.Param("device")

	data, status = models.GetDeviceByID(DeviceID)

	if status.Code == http.StatusOK {
		c.JSON(status.Code, gin.H{
			"msg":  status.Msg,
			"data": data,
		})
	} else {
		c.JSON(status.Code, gin.H{"msg": status.Msg})
	}
}

// GetDevices godoc
// @Summary get all devices
// @Tags Device
// @Accept  json
// @Success 200 {array} Device
// @Failure 400 {string} string "error"
// @Router /api/devices [GET]
func GetDevices(c *gin.Context) {
	var devices []models.Device
	var status = e.DefaultOk()

	auth_user, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := auth_user.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	devices, status = models.GetDevices(&user)

	c.JSON(status.Code, gin.H{
		"msg":  status.Msg,
		"data": devices,
	})
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
