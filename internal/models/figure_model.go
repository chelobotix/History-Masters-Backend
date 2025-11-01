package models

import (
	"gorm.io/gorm"
)

type FigureModelInput struct {
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	YearOfBirth    uint     `json:"year_of_birth"`
	YearOfDeath    uint     `json:"year_of_death"`
	CountryISOCode string   `json:"country_iso_code"`
	HistoricalEra  string   `json:"historical_era"`
	Areas          []string `json:"areas"`
	Profession     string   `json:"profession"`
	Achievements   []string `json:"achievements"`
}
type Figure struct {
	gorm.Model
	FirstName       string        `gorm:"not null;uniqueIndex:idx_firstname_lastname" validate:"required" json:"first_name"`
	LastName        string        `gorm:"not null;uniqueIndex:idx_firstname_lastname" validate:"required" json:"last_name"`
	YearOfBirth     Year          `gorm:"foreignKey:YearOfBirthID;" json:"year_of_birth"`
	YearOfBirthID   uint          `json:"-"`
	YearOfDeath     Year          `gorm:"foreignKey:YearOfDeathID;" json:"year_of_death"`
	YearOfDeathID   uint          `json:"-"`
	Country         Country       `gorm:"foreignKey:CountryID;preload:true" json:"country"`
	CountryID       uint          `json:"-"`
	HistoricalEra   HistoricalEra `gorm:"foreignKey:HistoricalEraID;" json:"historical_era"`
	HistoricalEraID uint          `json:"-"`
	Areas           []Area        `gorm:"many2many:figure_areas;"`
	Profession      Profession    `gorm:"foreignKey:ProfessionID;" json:"profession"`
	ProfessionID    uint          `json:"-"`
	Achievements    []Achievement `gorm:"many2many:figure_achievements;"`
}
