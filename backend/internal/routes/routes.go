package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Geovannisouza23/go-user-api/internal/handlers"
	"github.com/Geovannisouza23/go-user-api/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	protected := r.Group("/")
	protected.Use(middlewares.JWTAuth())
{
	protected.GET("/me", handlers.Me)

	protected.GET("/users", handlers.ListUsers)
	protected.GET("/users/:id", handlers.GetUser)
	protected.PUT("/users/:id", handlers.UpdateUser)
	protected.DELETE("/users/:id", handlers.DeleteUser)
}
}
