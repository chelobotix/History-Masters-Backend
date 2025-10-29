package models

import "gorm.io/gorm"

type Year struct {
	gorm.Model
	Year int  `gorm:"not null;uniqueIndex:idx_year_bc" validate:"required" json:"year"`
	BC   bool `gorm:"not null;default:false;uniqueIndex:idx_year_bc" json:"bc"`
}
