package internal

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

var bet_id = 1

func BetbyUserID(c *gin.Context) {
	name := c.Param("username")
	OppUser, mess := helpers.FindUser(name)
	if mess == "User not found" {
		c.JSON(404, gin.H{"message": mess})
		return
	}
	var betReq models.BetReq
	if err := c.BindJSON(&betReq); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	curr_ID, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(400, gin.H{"error": "User Cookie not found"})
		return
	}
	curr_PID, err := strconv.Atoi(curr_ID)
	if err != nil {
		c.JSON(400, gin.H{"error": "User Cookie not found"})
		return
	}
	opp_PID := OppUser.ID
	result := postgres.DB.Create(&models.Bet{
		BetID:            uint(bet_id),
		Name:             betReq.Name,
		Points:           float64(betReq.Points),
		Current_PlayerID: uint(curr_PID),
		Opp_PlayerID:     opp_PID,
		Activity_Start:   betReq.Activity_Start,
		Activity_End:     betReq.Activity_End,
	})
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Error in creating BET"})
		return
	}
	var curr models.User
	postgres.DB.First(&curr, curr_PID)

	OppUser.Points = OppUser.Points - float64(betReq.Points)
	curr.Points = curr.Points - float64(betReq.Points)
	postgres.DB.Save(&OppUser)
	postgres.DB.Save(&curr)
	bet_id++
	c.JSON(201, gin.H{"message": "Successfully created the BET"})
}
