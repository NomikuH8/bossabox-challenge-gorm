package models

import "gorm.io/gorm"

type ToolTag struct {
	gorm.Model
	TagName string
	ToolId  int
}

func (tt ToolTag) TableName() string {
	return "tool_tag"
}
