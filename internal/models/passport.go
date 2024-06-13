package models

type Passport struct {
	ID     int    `gorm:"passport_id"`
	Type   string `json:"passport_type"`
	Number string `json:"passport_number"`
}
