package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Points     float64 `json:"points"`
	Wins       int64   `json:"wins"`
	Bets       []*Bet  `gorm:"many2many:user_bets;"`
	HealthInfo HealthInfo
}

type UserSignin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Bet struct {
	gorm.Model
	BetID            uint    `json:"bet_id"`
	Name             string  `json:"name"`
	Points           float64 `json:"points"`
	Current_PlayerID uint    `json:"current_playerID"`
	Opp_PlayerID     uint    `json:"opp_playerID"`
	Activity_Start   string  `json:"activity_start"`
	Activity_End     string  `json:"activity_end"`
	Winner           string  `json:"winner"`
}

type BetReq struct {
	Name           string `json:"name"`
	Points         int64  `json:"points"`
	Activity_Start string `json:"activity_start"`
	Activity_End   string `json:"activity_end"`
}

type HealthInfo struct {
	gorm.Model
	UserID       uint  `json:"user_id"`
	Pedometer    int64 `json:"pedometer"`
	CalorieBurnt int64 `json:"calorie_burnt"`
	WaterCount   int64 `json:"water_count"`
	HeartRate    int64 `json:"heart_rate"`
}

type HealthInfoReq struct {
	Pedometer    int64 `json:"pedometer"`
	CalorieBurnt int64 `json:"calorie_burnt"`
	WaterCount   int64 `json:"water_count"`
	HeartRate    int64 `json:"heart_rate"`
}
