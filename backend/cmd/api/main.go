package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Geovannisouza23/go-user-api/internal/config"
	"github.com/Geovannisouza23/go-user-api/internal/database"
	"github.com/Geovannisouza23/go-user-api/internal/routes"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg)
	database.Migrate()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	routes.RegisterRoutes(r) // 👈 ESSA LINHA É O MAIS IMPORTANTE

	r.Run(":" + cfg.AppPort)
}
