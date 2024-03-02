package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/database"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/models"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	g := gin.Default()

	db := database.GetDatabase()
	err = db.AutoMigrate(&models.Tool{}, &models.ToolTag{})
	if err != nil {
		log.Panic(err)
	}

	routes.ToolRoutes(g)

	g.Run(":3000")
}
