package internal

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

var user_id = 1

func SSO(c *gin.Context) {
	var SignInReq models.UserSignin
	if err := c.BindJSON(&SignInReq); err != nil {
		c.JSON(500, gin.H{"error": "Binding JSON"})
		return
	}

	emailVer, lol := helpers.VerifyEmail(SignInReq.Email)
	if !emailVer {
		result := postgres.DB.Create(&models.User{
			Name:   SignInReq.Name,
			Email:  SignInReq.Email,
			Points: 100,
			UserID: uint(user_id),
		})
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Error in Signin User"})
			return
		}
		max_age := time.Now().UTC().Add(1 * time.Hour)
		c.SetCookie("user_id", strconv.FormatUint(uint64(user_id), 10), int(max_age.Unix()), "/", "", true, true)
		c.Request.Header.Add("user_id", strconv.FormatUint(uint64(user_id), 10))
		user_id++
		c.JSON(201, gin.H{
			"name":    SignInReq.Name,
			"email":   SignInReq.Email,
			"user_id": user_id - 1,
		})
	} else {
		max_age := time.Now().UTC().Add(1 * time.Hour)
		c.SetCookie("user_id", strconv.FormatUint(uint64(lol.UserID), 10), int(max_age.Unix()), "/", "", true, true)
		c.Request.Header.Add("user_id", strconv.FormatUint(uint64(lol.UserID), 10))
		c.JSON(200, gin.H{
			"user_id": lol.UserID,
		})
		return
	}
}
