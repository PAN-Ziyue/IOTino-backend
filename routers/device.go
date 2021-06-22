package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	status := e.DefaultOk()

	authUser, exist := c.Get("auth")

	if !exist {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)

	if !ok {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	deviceCount, onlineCount, dataCount := models.CountDevice(&user)
	chartData := models.GetChartData(&user)
	locationData := models.GetLocationData(&user)

	c.JSON(status.Code, gin.H{
		"msg":          status.Msg,
		"total":        deviceCount,
		"online":       onlineCount,
		"count":        dataCount,
		"chartData":    chartData,
		"locationData": locationData,
	})
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
	type DeviceJSON struct {
		Device string `json:"device"`
		Name   string `json:"name"`
	}

	var deviceJSON DeviceJSON
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

	authUser, exist := c.Get("auth")

	if !exist {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)

	if !ok {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	device.UserID = user.ID
	device.User = user

	status = models.CreateDevice(&device)

	c.JSON(status.Code, gin.H{"msg": status.Msg})
}

// GetDevices godoc
// @Summary get all devices
// @Tags Device
// @Accept  json
// @Success 200 {array} Device
// @Failure 400 {string} string "error"
// @Router /api/devices [GET]
func GetDevices(c *gin.Context) {
	var status = e.DefaultOk()

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	devices := models.GetDevices(&user)

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
	type DeviceJSON struct {
		Device string `json:"device"`
		Name   string `json:"name"`
	}

	var deviceJSON DeviceJSON
	var status = e.DefaultOk()

	// bind model
	err := c.BindJSON(&deviceJSON)
	if err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// get the device
	device, err := models.GetDeviceByID(deviceJSON.Device)

	if err != nil {
		status.Set(http.StatusBadRequest, e.DeviceNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	device.Name = deviceJSON.Name

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusUnauthorized, e.UserNotFound)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	status = models.UpdateDevice(&user, &device)

	c.JSON(status.Code, gin.H{"msg": status.Msg})
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
	type DeviceJSON struct {
		Device string `json:"device"`
	}
	var deviceJSON DeviceJSON
	var status = e.DefaultOk()

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	err := c.BindJSON(&deviceJSON)
	if err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	err = models.DeleteDevice(&user, deviceJSON.Device)
	if err != nil {
		status.Set(http.StatusBadRequest, e.CannotDeleteDevice)
	}

	c.JSON(status.Code, gin.H{"msg": status.Msg})
}
