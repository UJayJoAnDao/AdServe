package model

import (
	"time"

	"gorm.io/gorm"
)

type Ad struct {
	Id         uint        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title      string      `json:"title" binding:"required"`
	StartAt    time.Time   `json:"startAt" binding:"required"`
	EndAt      time.Time   `json:"endAt" binding:"required"`
	Conditions AdCondition `json:"conditions"`
	gorm.Model
}

type AdCondition struct {
	AgeStart int      `json:"ageStart"`
	AgeEnd   int      `json:"ageEnd"`
	Gender   string   `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}
