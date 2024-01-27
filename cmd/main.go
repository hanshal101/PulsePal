package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/ppSM/database/postgres"
	"github.com/hanshal101/ppSM/internal"
)

func init() {
	postgres.PostgresInitializer()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", internal.Get)
	r.POST("/signin", internal.SSO)
	r.POST("/users/:username", internal.FindbyName)
	r.POST("/users/:username/bet", internal.BetbyUserID)
	r.GET("/profile/bets", internal.GetUsersBET)
	r.POST("/profile/uploadHealth", internal.UploadHealthData)
	r.POST("/profile/bets/:betID", internal.EndBet)
	r.GET("/leadboard/all", internal.WorldLeadBoard)
	r.GET("/users", internal.AllUsers)
	r.POST("/profile", internal.ProfileData)
	r.Run(":9876")
}
