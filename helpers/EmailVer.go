package helpers

import (
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func VerifyEmail(userEmail string) (bool, models.User) {
	var user models.User
	result := postgres.DB.Where(&models.User{Email: userEmail}).First(&user)
	if result.Error != nil {
		return false, user
	}
	return true, user
}
