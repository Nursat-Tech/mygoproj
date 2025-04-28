package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/db"
	"github.com/nursat/myproj/internal/routes"
)

func main() {
	db.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
