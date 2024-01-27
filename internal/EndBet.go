package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

func EndBet(c *gin.Context) {
	var result string
	bet_id := c.Param("betID")
	BID, err := helpers.STRtoUINT(bet_id)
	if err != nil {
		c.JSON(400, gin.H{"Error": "Error in betID"})
		return
	}

	var bet models.Bet
	var user1_health models.HealthInfo
	var user2_health models.HealthInfo
	var user1 models.User
	var user2 models.User

	// Load the bet
	postgres.DB.First(&bet, BID)

	if bet.Winner != "" {
		c.JSON(400, gin.H{"Error": "The BET has already completed"})
		return
	}

	// Load the users associated with the bet
	postgres.DB.First(&user1_health, bet.Current_PlayerID)
	postgres.DB.First(&user2_health, bet.Opp_PlayerID)

	postgres.DB.First(&user1, user1_health.UserID)
	postgres.DB.First(&user2, user2_health.UserID)

	print("User1 CC : ", user1_health.CalorieBurnt)
	print("User2 CC : ", user2_health.CalorieBurnt)

	// Compare HealthInfo
	if user1_health.CalorieBurnt > user2_health.CalorieBurnt {
		user1.Points += 2 * bet.Points
		user1.Wins++
		postgres.DB.Save(&user1)
		result = user1.Name
		bet.Winner = string(rune(user1.UserID))
	} else if user1.HealthInfo.CalorieBurnt < user2.HealthInfo.CalorieBurnt {
		user2.Points += 2 * bet.Points
		user2.Wins++
		postgres.DB.Save(&user2)
		result = user2.Name
		bet.Winner = string(rune(user2.UserID))
	} else {
		user1.Points += bet.Points
		user1.Wins++
		postgres.DB.Save(&user1)
		user2.Points += bet.Points
		user2.Wins++
		postgres.DB.Save(&user2)
		result = "Draw"
	}
	bet.Winner = result
	postgres.DB.Save(&bet)

	c.JSON(200, gin.H{
		"bet_id":  BID,
		"message": "Completed Successfully",
		"winner":  result,
	})
}
