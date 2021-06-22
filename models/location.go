package models

type Location struct {
	ID        uint    `json:"-" gorm:"primaryKey" swaggerignore:"true"`
	DeviceID  uint    `json:"-" swaggerignore:"true"`
	Device    Device  `json:"-" gorm:"foreignKey:DeviceID" swaggerignore:"true"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Time      int64   `json:"time"`
}
