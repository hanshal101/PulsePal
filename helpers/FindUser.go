package helpers

import (
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func FindUser(username string) (models.User, string) {
	var user models.User
	result := postgres.DB.Where(&models.User{Name: username}).First(&user)
	if result.Error != nil {
		return user, "User not found"
	}
	return user, ""
}

func FindUserbyID(id uint) (models.User, string) {
	var user models.User
	result := postgres.DB.Where(&models.User{UserID: id}).First(&user)
	if result.Error != nil {
		return user, "User not found"
	}
	return user, ""
}
