package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Owner    int    `gorm:"not null" json:"owner_id"`
	GenderID int    `gorm:"not null" json:"gender_id"`
	RealmID  int    `gorm:"not null" json:"realm_id"`
}
