package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/auth"
	"github.com/nursat/myproj/internal/db"
	"github.com/nursat/myproj/internal/handler"
	"github.com/nursat/myproj/internal/middleware"
	"github.com/nursat/myproj/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	playerService := service.NewPlayerService(db.DB)
	playerHandler := handler.NewPlayerHandler(playerService)

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me) // Защищённый эндпоинт
	}

	players := r.Group("/players", middleware.AuthRequired())
	{
		players.POST("/", playerHandler.CreatePlayer)
		players.GET("/", playerHandler.GetPlayers)
		players.GET("/:id", playerHandler.GetPlayerByID)
		players.PUT("/:id", playerHandler.UpdatePlayer)
		players.DELETE("/:id", playerHandler.DeletePlayer)
	}
}
