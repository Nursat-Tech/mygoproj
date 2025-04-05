package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/handler"
	"github.com/nursat/myproj/internal/models"
	"github.com/nursat/myproj/internal/repository"
	"github.com/nursat/myproj/internal/routes"
	"github.com/nursat/myproj/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "postgres://postgres:51177477@postgres:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = db.AutoMigrate(&models.Player{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	playerRepo := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerService)

	r := gin.Default()
	routes.RegisterPlayerRoutes(r, playerHandler)

	r.Run(":8080")
}
