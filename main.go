package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/routes"
)

func main() {
	g := gin.Default()

	routes.ToolRoutes(g)

	g.Run(":3000")
}
