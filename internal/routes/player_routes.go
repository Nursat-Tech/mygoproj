package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/handler"
)

func RegisterPlayerRoutes(r *gin.Engine, handler *handler.PlayerHandler) {
	playerRoutes := r.Group("/players")
	{
		playerRoutes.POST("/", handler.CreatePlayer)
		playerRoutes.GET("/", handler.GetPlayers)
		playerRoutes.GET("/:id", handler.GetPlayerByID)
		playerRoutes.PUT("/:id", handler.UpdatePlayer)
		playerRoutes.DELETE("/:id", handler.DeletePlayer)
	}
}
