package models

import "time"

type BookRecord struct {
	ReservationNum string    `json:"reservationId" gorm:"type:varchar(20);primary_key;not null"`
	PricePerHour   float64   `json:"pricePerHour" gorm:"type:real;not null"`
	PricePerDay    float64   `json:"pricePerDay" gorm:"type:real;not null"`
	ReservedDate   time.Time `json:"reservedDate" gorm:"type:timestamp with time zone;not null"`
	PickUpTime     time.Time `json:"pickUpTime" gorm:"type:timestamp with time zone"`
	DropOfTime     time.Time `json:"dropOfTime" gorm:"type:timestamp with time zone"`
	CreateAt       time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdateAt       time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
	CarId          int       `json:"carId" gorm:"type:integer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Car            Car       `json:"car" gorm:"foreignKey:CarId"`
	LocationId     int       `json:"locationId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Location       Location  `json:"location" gorm:"foreignKey:LocationId"`
	CustomerId     string    `json:"customerId" gorm:"type:varchar(30);not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Customer       Customer  `json:"customer" gorm:"foreignKey:CustomerId;"`
}
