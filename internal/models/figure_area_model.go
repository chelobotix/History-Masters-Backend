package models

type FigureAreas struct {
	FigureID uint `gorm:"primaryKey"`
	AreaID   uint `gorm:"primaryKey"`
}
