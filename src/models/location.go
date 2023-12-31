package models

import "time"

type Location struct {
	Id            int       `json:"locationId" gorm:"primary_key"`
	StreetAddress string    `json:"streetAddress" gorm:"type:varchar(100)"`
	Telephone     string    `json:"telephone" gorm:"type:varchar(30)"`
	CreateAt      time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdateAt      time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
}
