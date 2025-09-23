package api

import (
	"github.com/gin-gonic/gin"
	"rest.com/m/internal/api/handlers"
	"rest.com/m/internal/middleware"
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
