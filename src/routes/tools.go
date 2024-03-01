package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nomikuh8/bossabox-challenge-gorm/src/controllers"
)

func ToolRoutes(g *gin.Engine) {
	t := g.Group("/api/v1/tools")
	{
		tc := controllers.NewToolsController()
		t.GET("/", tc.GetAll)
	}
}
