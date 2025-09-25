package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ritu-p/go-rest/internal/api/handlers"
	"github.com/ritu-p/go-rest/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// global middleware
	r.Use(middleware.RequestLogger())

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", handlers.CreateUser)
			users.GET("/:id", handlers.GetUser)
			users.GET("/", handlers.ListUsers)
		}
	}

	return r
}
