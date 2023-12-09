package models

import "time"

type Customer struct {
	Uid        string    `json:"uid" gorm:"type:varchar(255);primary_key"`
	Firstname  string    `json:"firstname" gorm:"type:varchar(255); not null"`
	Lastname   string    `json:"lastname" gorm:"type:varchar(255); not null"`
	Address    string    `json:"address" gorm:"type:varchar(255); not null"`
	Email      string    `json:"email" gorm:"type:varchar(255); not null"`
	CreditCard string    `json:"creditCard" gorm:"type:varchar(255); not null"`
	IsStudent  *bool     `json:"isStudent" gorm:"type:boolean; not null"`
	Telephone  string    `json:"telephone" gorm:"type:varchar(255); not null"`
	Phone      string    `json:"phone" gorm:"type:varchar(255); not null"`
	Licence    int       `json:"licence" gorm:"type:integer;not null"`
	Tickets    int       `json:"tickets" gorm:"type:integer;not null"`
	StateIssue string    `json:"stateIssue" gorm:"type:enum_customers_stateissue;not null"`
	Expiration time.Time `json:"expiration" gorm:"type:timestamp with time zone"`
	CreateAt   time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoCreateTime"`
	UpdateAt   time.Time `json:"-" gorm:"type:timestamp with time zone;not null;autoUpdateTime"`
}
