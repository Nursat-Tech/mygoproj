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
		protected.GET("/me", auth.Me)
	}

	admin := r.Group("/players")
	admin.Use(middleware.AuthRequired(), middleware.RequireRole("admin"))
	{
		admin.POST("/", playerHandler.CreatePlayer)
		admin.PUT("/:id", playerHandler.UpdatePlayer)
		admin.DELETE("/:id", playerHandler.DeletePlayer)
	}

	players := r.Group("/players")
	players.Use(middleware.AuthRequired())
	{
		players.GET("/", playerHandler.GetPlayers)
		players.GET("/:id", playerHandler.GetPlayerByID)
	}
}
