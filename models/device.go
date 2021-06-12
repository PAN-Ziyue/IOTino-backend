package models

import (
    "IOTino/pkg/e"
    "log"
    "net/http"
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
    ID               uint    `json:"-" gorm:"primaryKey" swaggerignore:"true"`
    UserID           uint    `json:"-" swaggerignore:"true"`
    User             User    `json:"-" gorm:"foreignKey:UserID" swaggerignore:"true"`
    Device           string  `json:"device" gorm:"unique;size:255"`
    Name             string  `json:"name" gorm:"unique;size:255"`
    Alert            bool    `json:"alert"`
    Count            uint64  `json:"count" gorm:"default:0"`
    CurrentLatitude  float64 `json:"current_latitude"`
    CurrentLongitude float64 `json:"current_longitude"`
    Value            int64   `json:"value" gorm:"default:0"`
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
func GetDeviceByID(DeviceID string) (Device, e.Status) {
    var device Device
    status := e.DefaultOk()

    result := DB.Where("Device = ?", DeviceID).First(&device)

    if result.RowsAffected == 0 {
        status.Set(http.StatusOK, e.DeviceNotFound)
    }

    return device, status
}

func GetDevices(user *User) ([]Device, e.Status) {
    var devices []Device
    status := e.DefaultOk()

    result := DB.Where("user_id = ?", user.ID).Find(&devices)

    if result.RowsAffected == 0 {
        status.Set(http.StatusOK, e.NoDevices)
    }

    return devices, status
}

func DeleteDevice(user *User, deviceID string) e.Status {
    device, status := GetDeviceByID(deviceID)

    err := DB.Where("user_id = ?", user.ID).Delete(&device).Error

    if err != nil {
        status.Set(http.StatusNotFound, e.CannotDeleteDevice)
    }

    return status
}

func UpdateDevice(user *User, newDevice *Device) e.Status {
    status := e.DefaultOk()

    err := DB.Where("user_id = ?", user.ID).Save(newDevice).Error

    if err != nil {
        status.Set(http.StatusNotFound, e.CannotUpdateDevice)
    }

    return status
}

func HandleMQTT(msg *MQTTMsg) {
    log.Println("An MQTT message is processed")

    // device
    var device Device

    // get device
    err := DB.Where("device = ?", msg.ClientID).First(&device).Error
    if err != nil {
        log.Println("Invalid MQTT message, due to", err)
        return
    }

    device.Alert = msg.Alert == 0
    device.CurrentLongitude = msg.Longitude
    device.CurrentLatitude = msg.Latitude
    device.Count++
    device.Value = msg.Value

    DB.Save(&device)
}
