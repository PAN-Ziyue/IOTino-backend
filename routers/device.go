package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"net/http"

	"github.com/astaxie/beego/validation"
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
	var data models.Device
	var status = e.DefaultStatus()
	var err error

	// bind model
	err = c.BindJSON(&data)

	if err != nil {
		status.Set(http.StatusBadRequest, e.BadJson)
	} else {
		valid := validation.Validation{}
		if !valid.HasErrors() {
			status = models.CreateDevice(&data)
		} else {
			status.Set(http.StatusBadRequest, e.BadParameter)
		}
	}

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
	var status = e.DefaultStatus()
	var DeviceID string = c.Param("device")

	status, data = models.GetDeviceByID(DeviceID)

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
	//var data []models.Device


}





