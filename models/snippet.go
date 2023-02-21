package models

import "github.com/jinzhu/gorm"

type Snippet struct {
	gorm.Model
	Title        string `json:"title"`
	Code         string `json:"code"`
	IsActive     bool   `json:"isActive" gorm:"default:true"`
	PlatformType int    `json:"platformType"`
}
