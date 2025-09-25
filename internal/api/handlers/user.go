package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritu-p/go-rest/internal/db"
	"github.com/ritu-p/go-rest/internal/models"
)

// CreateUser -> POST /api/v1/users
func CreateUser(c *gin.Context) {
	var payload struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{Name: payload.Name, Email: payload.Email}
	if err := db.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, u)
}

// GetUser -> GET /api/v1/users/:id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	if err := db.DB.First(&u, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

// ListUsers -> GET /api/v1/users
func ListUsers(c *gin.Context) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
