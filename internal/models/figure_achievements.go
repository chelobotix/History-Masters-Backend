package models

type FigureAchievements struct {
	FigureID      uint `gorm:"primaryKey"`
	AchievementID uint `gorm:"primaryKey"`
}
