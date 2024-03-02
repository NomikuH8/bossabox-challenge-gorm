package models

import "gorm.io/gorm"

type ToolSend struct {
	Id          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Link        string   `json:"link" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
}

type Tool struct {
	gorm.Model
	Title       string
	Link        string
	Description string
	Tags        []ToolTag
}

func (t Tool) TableName() string {
	return "tool"
}
