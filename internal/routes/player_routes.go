package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/db"
	"github.com/nursat/myproj/internal/handler"
	"github.com/nursat/myproj/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	playerService := service.NewPlayerService(db.DB)
	playerHandler := handler.NewPlayerHandler(playerService)

	players := r.Group("/players")
	{
		players.POST("/", playerHandler.CreatePlayer)
		players.GET("/", playerHandler.GetPlayers)
		players.GET("/:id", playerHandler.GetPlayerByID)
		players.PUT("/:id", playerHandler.UpdatePlayer)
		players.DELETE("/:id", playerHandler.DeletePlayer)
	}
}
