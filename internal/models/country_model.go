package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name         string `gorm:"unique;not null" validate:"required" json:"name"`
	NameEs       string `json:"name_es"`
	ActualName   string `json:"actual_name"`
	ActualNameEs string `json:"actual_name_es"`
	IsoCode      string `gorm:"unique;" json:"iso_code"`
}
