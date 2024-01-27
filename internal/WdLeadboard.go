package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func WorldLeadBoard(c *gin.Context) {
	var data []models.HealthInfo

	postgres.DB.Order("calorie_burnt DESC").Find(&data)

	c.JSON(200, data)
}
