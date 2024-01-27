package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/helpers"
	"github.com/hanshal101/ppSM/models"
)

func UploadHealthData(c *gin.Context) {
	var HealthDataReq models.HealthInfoReq
	var User models.User
	var UserHealth models.HealthInfo

	if err := c.BindJSON(&HealthDataReq); err != nil {
		c.JSON(400, gin.H{"Error": "Error in Binding"})
		return
	}

	userCID, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(400, gin.H{"Error": "User not recognized, go to /signin page"})
		return
	}
	userPID, err := helpers.STRtoUINT(userCID)
	if userPID == 0 || err != nil {
		c.JSON(400, gin.H{"Error": "User not recognized, go to /signin page"})
		return
	}

	// Use First instead of Find
	postgres.DB.First(&User, userPID)

	// Check if user is found
	if User.ID == 0 {
		c.JSON(404, gin.H{"Error": "User not found"})
		return
	}

	ver := postgres.DB.First(&UserHealth, userPID)
	if ver.Error != nil {
		User.HealthInfo.Pedometer = HealthDataReq.Pedometer
		User.HealthInfo.CalorieBurnt = HealthDataReq.CalorieBurnt
		User.HealthInfo.WaterCount = HealthDataReq.WaterCount
		User.HealthInfo.HeartRate = HealthDataReq.HeartRate

		// Save the changes to the existing HealthInfo record
		postgres.DB.Create(&models.HealthInfo{
			Pedometer:    HealthDataReq.Pedometer,
			CalorieBurnt: HealthDataReq.CalorieBurnt,
			WaterCount:   HealthDataReq.WaterCount,
			HeartRate:    HealthDataReq.HeartRate,
			UserID:       userPID,
		})

		c.JSON(201, gin.H{
			"user_id":       userPID,
			"health_id":     User.HealthInfo.ID,
			"pedometer":     User.HealthInfo.Pedometer,
			"calorie_burnt": User.HealthInfo.CalorieBurnt,
			"water_count":   User.HealthInfo.WaterCount,
			"heart_rate":    User.HealthInfo.HeartRate,
		})
		return
	}
	User.HealthInfo = UserHealth

	User.HealthInfo.Pedometer = HealthDataReq.Pedometer
	User.HealthInfo.CalorieBurnt = HealthDataReq.CalorieBurnt
	User.HealthInfo.WaterCount = HealthDataReq.WaterCount
	User.HealthInfo.HeartRate = HealthDataReq.HeartRate

	// Save the changes to the existing HealthInfo record
	postgres.DB.Save(&User.HealthInfo)

	c.JSON(201, gin.H{
		"user_id":       userPID,
		"health_id":     User.HealthInfo.ID,
		"pedometer":     User.HealthInfo.Pedometer,
		"calorie_burnt": User.HealthInfo.CalorieBurnt,
		"water_count":   User.HealthInfo.WaterCount,
		"heart_rate":    User.HealthInfo.HeartRate,
	})

}
