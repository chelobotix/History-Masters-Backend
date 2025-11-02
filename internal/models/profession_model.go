package models

import "gorm.io/gorm"

type Profession struct {
	gorm.Model
	Name          string   `gorm:"unique;not null" validate:"required" json:"name"`
	NameEs        string   `json:"name_es"`
	Description   string   `json:"description"`
	DescriptionEs string   `json:"description_es"`
	Figures       []Figure `gorm:"many2many:figure_professions;" json:"-"`
}
