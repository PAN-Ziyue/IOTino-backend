package models

import (
	"IOTino/pkg/e"
	"log"
	"net/http"
	"sort"
)

type MQTTMsg struct {
	Alert     int     `json:"alert"`
	ClientID  string  `json:"clientId"`
	Info      string  `json:"info"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Timestamp int64   `json:"timestamp"`
	Value     int64   `json:"value"`
}

type Device struct {
	ID               uint       `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	UserID           uint       `json:"-" swaggerignore:"true"`
	User             User       `json:"-" gorm:"foreignKey:UserID" swaggerignore:"true"`
	Device           string     `json:"device" gorm:"unique;size:255"`
	Name             string     `json:"name" gorm:"unique;size:255"`
	Count            int64      `json:"count" gorm:"default:0"`
	CurrentLatitude  float64    `json:"current_latitude"`
	CurrentLongitude float64    `json:"current_longitude"`
	Value            int64      `json:"value" gorm:"default:0"`
	Status           string     `json:"status" gorm:"default:'offline'"`
	Trace            []Location `json:"trace" gorm:"-"`
}

type ChartData struct {
	X string `json:"x"`
	Y int64  `json:"y"`
}

type LocationData struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type ChartDataSlice []ChartData

const (
	offline = "offline"
	normal  = "normal"
	alert   = "alert"
)

// CountDevice
// Get the total number of the devices
func CountDevice(user *User) (int64, int64, int64) {
	var deviceCount int64 = 0
	var onlineCount int64 = 0
	var dataCount int64 = 0

	DB.Model(&Device{}).Where("user_id = ?", user.ID).Count(&deviceCount)
	DB.Model(&Device{}).Where("user_id = ? AND status <> ?", user.ID, offline).Count(&onlineCount)

	var devices []Device

	DB.Model(&Device{}).Where("user_id = ?", user.ID).Find(&devices)
	for _, d := range devices {
		dataCount += d.Count
	}

	return deviceCount, onlineCount, dataCount
}

//>>>>>>>>>>>>>>
// sort count
//<<<<<<<<<<<<<<

func (a ChartDataSlice) Len() int {
	return len(a)
}
func (a ChartDataSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ChartDataSlice) Less(i, j int) bool {
	return a[j].Y < a[i].Y
}

func GetChartData(user *User) []ChartData {
	var chartData []ChartData
	var devices []Device

	DB.Model(&Device{}).Where("user_id = ?", user.ID).Find(&devices)
	for _, d := range devices {
		chartData = append(chartData, ChartData{
			X: d.Name,
			Y: d.Count,
		})
	}

	sort.Sort(ChartDataSlice(chartData))
	return chartData
}

func GetLocationData(user *User) []LocationData {
	var locationData []LocationData
	var devices []Device

	DB.Model(&Device{}).Where("user_id = ?", user.ID).Find(&devices)
	for _, d := range devices {
		locationData = append(locationData, LocationData{
			Name:      d.Name,
			Longitude: d.CurrentLongitude,
			Latitude:  d.CurrentLatitude,
		})
	}

	return locationData
}

// CreateDevice
// Create a device
func CreateDevice(device *Device) e.Status {
	var DuplicateDevices []Device
	status := e.New(http.StatusCreated, e.DeviceCreated)

	resultDevice := DB.Where("Device = ?", device.Device).Find(&DuplicateDevices)
	resultName := DB.Where("Name = ?", device.Name).Find(&DuplicateDevices)

	if resultDevice.RowsAffected > 0 || resultName.RowsAffected > 0 {
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
func GetDeviceByID(DeviceID string) (Device, error) {
	var device Device

	err := DB.Where("Device = ?", DeviceID).First(&device).Error
	if err != nil {
		return device, nil
	}

	return device, nil
}

func GetDevices(user *User) []Device {
	var devices []Device

	DB.Where("user_id = ?", user.ID).Find(&devices)

	for i := range devices {
		var trace []Location
		DB.Order("time").Where("device_id = ?", devices[i].ID).Limit(10).Find(&trace)
		devices[i].Trace = trace
	}

	return devices
}

func DeleteDevice(user *User, deviceID string) error {
	device, err := GetDeviceByID(deviceID)

	if err != nil {
		return err
	}

	err = DB.Where("user_id = ?", user.ID).Delete(&device).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateDevice(user *User, newDevice *Device) e.Status {
	status := e.DefaultOk()

	err := DB.Where("user_id = ?", user.ID).Save(newDevice).Error

	if err != nil {
		status.Set(http.StatusConflict, e.CannotUpdateDevice)
	}

	return status
}

func HandleMQTT(msg *MQTTMsg) {
	log.Println("An MQTT message is processed")

	// device
	var device Device
	var location Location

	// get device
	err := DB.Where("device = ?", msg.ClientID).First(&device).Error
	if err != nil {
		log.Println("Invalid MQTT message, due to", err)
		return
	}

	if msg.Alert == 0 {
		device.Status = normal
	} else {
		device.Status = alert
	}

	device.CurrentLongitude = msg.Longitude
	device.CurrentLatitude = msg.Latitude
	device.Count++
	device.Value = msg.Value

	DB.Save(&device)

	location.Device = device
	location.DeviceID = device.ID
	location.Latitude = device.CurrentLatitude
	location.Longitude = device.CurrentLongitude
	location.Time = msg.Timestamp

	DB.Create(&location)
}
