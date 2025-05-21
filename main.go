package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/db"
	"github.com/nursat/myproj/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.Static("/static", "./footsite/static")

	r.GET("/", func(c *gin.Context) {
		c.File("./footsite/home.html")
	})

	r.GET("/login", func(c *gin.Context) {
		c.File("./footsite/login.html")
	})

	r.GET("/register", func(c *gin.Context) {
		c.File("./footsite/register.html")
	})

	routes.SetupRoutes(r) // ✅ Дұрыс

	r.Run(":8080")
}
