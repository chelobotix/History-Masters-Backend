package models

import "gorm.io/gorm"

type HistoricalEra struct {
	gorm.Model
	Name   string `gorm:"unique;not null" validate:"required" json:"name"`
	NameEs string `json:"name_es"`
}
