package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func AllUsers(c *gin.Context) {
	var data []models.User

	err := postgres.DB.Find(&data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, data)
}
