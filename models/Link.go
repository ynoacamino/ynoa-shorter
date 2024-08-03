package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model

	Short  string `gorm:"primaryKey" json:"short"`
	Real   string `gorm:"not null" json:"real"`
	UserId string `json:"userId"`
	Public bool   `gorm:"not null" json:"public"`
}
