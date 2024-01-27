package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

func FindbyName(c *gin.Context) {
	name := c.Param("username")
	User, mess := helpers.FindUser(name)
	if mess == "User not found" {
		c.JSON(404, gin.H{"message": mess})
		return
	}

	// curr_ID, err := c.Cookie("user_id")
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "User Cookie not found"})
	// 	return
	// }

	// curr_PID, err := strconv.Atoi(curr_ID)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid user ID in cookie"})
	// 	return
	// }

	var bets []models.Bet

	err := postgres.DB.Where("current_player_id = ? OR opp_player_id = ?", User.ID, User.ID).Find(&bets).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching bets"})
		return
	}
	c.JSON(200, gin.H{
		"user_id":   User.ID,
		"name":      User.Name,
		"email":     User.Email,
		"points":    User.Points,
		"user_bets": bets,
	})
}
