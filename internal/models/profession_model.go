package models

import "gorm.io/gorm"

type Profession struct {
	gorm.Model
	Name string `gorm:"unique;not null" validate:"required" json:"name"`
}
