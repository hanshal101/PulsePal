package main

import (
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/models"
)

func init() {
	postgres.PostgresInitializer()
}

func main() {
	postgres.DB.AutoMigrate(&models.User{})
	postgres.DB.AutoMigrate(&models.Bet{})
	postgres.DB.AutoMigrate(&models.HealthInfo{})
}
