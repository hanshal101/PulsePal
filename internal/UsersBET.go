package internal

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func GetUsersBET(c *gin.Context) {
	curr_ID, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(400, gin.H{"error": "User Cookie not found"})
		return
	}

	curr_PID, err := strconv.Atoi(curr_ID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID in cookie"})
		return
	}

	var bets []models.Bet

	err = postgres.DB.Where("current_player_id = ? OR opp_player_id = ?", curr_PID, curr_PID).Find(&bets).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching bets"})
		return
	}

	c.JSON(200, bets)
}
