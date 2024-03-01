package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/database"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/models"
)

func NewToolsController() toolsController {
	return toolsController{}
}

type toolsController struct{}

func (tc *toolsController) GetAll(c *gin.Context) {
	db := database.GetDatabase()

	tools := []models.Tool{}
	db.Find(&tools)

	c.JSON(200, tools)
}
