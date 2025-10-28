package models

import (
	"gorm.io/gorm"
)

type Figure struct {
	gorm.Model
	FirstName       string        `gorm:"not null;uniqueIndex:idx_firstname_lastname" validate:"required" json:"first_name"`
	LastName        string        `gorm:"not null;uniqueIndex:idx_firstname_lastname" validate:"required" json:"last_name"`
	YearOfBirth     Year          `gorm:"foreignKey:YearOfBirthID;" json:"year_of_birth"`
	YearOfBirthID   uint          `json:"-"`
	Country         Country       `gorm:"foreignKey:CountryID;preload:true" json:"country"`
	CountryID       uint          `json:"-"`
	HistoricalEra   HistoricalEra `gorm:"foreignKey:HistoricalEraID;" json:"historical_era"`
	HistoricalEraID uint          `json:"-"`
	Areas           []Area        `gorm:"many2many:figure_areas;"`
	Profession      Profession    `gorm:"foreignKey:ProfessionID;" json:"profession"`
	ProfessionID    uint          `json:"-"`
	Achievements    []Achievement `gorm:"many2many:figure_achievements;"`
}
