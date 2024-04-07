package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Ad struct {
	Id         uint          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title      string        `json:"title" binding:"required"`
	StartAt    time.Time     `json:"startAt" binding:"required"`
	EndAt      time.Time     `json:"endAt" binding:"required"`
	Conditions []AdCondition `gorm:"foreignKey:AdID" json:"conditions"`
	gorm.Model
}

type AdCondition struct {
	ID       uint         `gorm:"primary_key;AUTO_INCREMENT" json:"-"` // 外鍵 json:"-" 表示不會被序列化
	AdID     uint         `gorm:"index"`                               // 外鍵 json:"-" 表示不會被序列化
	AgeStart *int         `json:"ageStart"`
	AgeEnd   *int         `json:"ageEnd"`
	Gender   *StringSlice `json:"gender" grom:"type:json"`
	Country  *StringSlice `json:"country" gorm:"type:json"`
	Platform *StringSlice `json:"platform" gorm:"type:json"`
}

type Result struct {
	Title string
	EndAt time.Time `json:"endAt"`
}

type Response struct {
	Items []Result `json:"items"`
}
type StringSlice []string

func (s *StringSlice) Scan(value interface{}) error { // Scan是GORM庫的一個方法，它將從數據庫中讀取值並將其掃描到結構體中，會自動被調用嗎？A: 是的，當我們從數據庫中讀取StringSlice類型的值時，GORM庫將自動調用這個方法。
	return json.Unmarshal(value.([]byte), s) // []string -> []byte
}

func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// func (c *AdCondition) setParams(ageStart, ageEnd *int, gender *string, country, platform []string) {
// 	c.AgeStart = ageStart
// 	c.AgeEnd = ageEnd
// 	c.Gender = gender
// 	c.Country = (*StringSlice)(&country)
// 	c.Platform = (*StringSlice)(&platform)
// }
