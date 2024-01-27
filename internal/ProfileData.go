package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func ProfileData(c *gin.Context) {
	var user models.User
	var user_id uint
	if err := c.BindJSON(&user_id); err != nil {
		c.JSON(400, gin.H{"message": "Error in userID"})
		return
	}

	postgres.DB.First(&user, user_id)
	c.JSON(200, gin.H{
		"name":   user.Name,
		"email":  user.Email,
		"points": user.Points,
		"wins":   user.Wins,
	})
}
