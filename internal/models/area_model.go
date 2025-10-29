package models

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	Name    string   `gorm:"unique;not null" validate:"required" json:"name"`
	Figures []Figure `gorm:"many2many:figure_areas;" json:"-"`
}
