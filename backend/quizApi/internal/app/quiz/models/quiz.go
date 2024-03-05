package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Active      bool       `gorm:"default:false"`
	StartTime   *time.Time `gorm:"type:timestamp"`
	EndTime     *time.Time `gorm:"type:timestamp"`
	Questions   []Question
}

type Question struct {
	gorm.Model
	QuizID   uint
	Question string `gorm:"not null"`
	Options  []Option `gorm:"foreignKey:QuestionID"`
}

type Option struct {
	gorm.Model
	Text       string `gorm:"not null"`
	IsCorrect  bool
	QuestionID uint
}
