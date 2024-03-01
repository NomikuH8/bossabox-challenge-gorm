package models

import "gorm.io/gorm"

type Tool struct {
	gorm.Model
	Title       string
	Link        string
	Description string
}

func (Tool) TableName() string {
	return "tool"
}
