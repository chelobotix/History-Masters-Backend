package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name    string `gorm:"unique;not null" validate:"required" json:"name"`
	IsoCode string `gorm:"unique;not null" validate:"required" json:"iso_code"`
}
