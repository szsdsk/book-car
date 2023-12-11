package models

import "time"

type Car struct {
	Id           int       `json:"carId" gorm:"primary_key"`
	Make         string    `json:"make" gorm:"type:varchar(20);not null"`
	Model        string    `json:"model" gorm:"type:varchar(20);not null"`
	PricePerHour float64   `json:"pricePerHour" gorm:"type:real;not null"`
	PricePerDay  float64   `json:"pricePerDay" gorm:"type:real;not null"`
	Capacity     int       `json:"capacity" gorm:"type:integer;not null"`
	Description  string    `json:"description" gorm:"type:varchar(255);not null"`
	CreateAt     time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdateAt     time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	Img          string    `json:"img" gorm:"varchar(100);not null"`
}
