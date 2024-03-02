package controllers

import (
	"log"

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
	err := db.Preload("Tags").Find(&tools).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error while trying to retrieve tools",
		})
		return
	}

	toolsSend := []models.ToolSend{}
	for _, tool := range tools {
		toolSend := models.ToolSend{
			Title:       tool.Title,
			Link:        tool.Link,
			Description: tool.Description,
		}

		toolSend.Tags = []string{}
		for _, tag := range tool.Tags {
			toolSend.Tags = append(toolSend.Tags, tag.TagName)
		}

		toolsSend = append(toolsSend, toolSend)
	}

	c.JSON(200, toolsSend)
}

func (tc *toolsController) InsertOne(c *gin.Context) {
	var body models.ToolSend

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "body invalid",
		})
		return
	}

	db := database.GetDatabase()

	toolInsert := models.Tool{
		Title:       body.Title,
		Link:        body.Link,
		Description: body.Description,
	}

	err = db.Create(&toolInsert).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while trying to insert tool",
		})
		log.Panic(err)
	}

	for _, tag := range body.Tags {
		toolTag := models.ToolTag{
			TagName: tag,
			ToolId:  int(toolInsert.ID),
		}

		err := db.Create(&toolTag).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error while trying to insert tool tag",
			})
		}
	}

	body.Id = int(toolInsert.ID)

	c.JSON(200, body)
}
