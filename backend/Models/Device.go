package Models

type Device struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
