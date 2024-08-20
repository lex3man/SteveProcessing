package models

import (
	"gorm.io/gorm"
)

type BotConfig struct {
	gorm.Model
	Caption string `json:"caption"`
	Token   string `json:"token"`
	Status  string `json:"status"`
	Active  bool   `json:"active"`
}

type AddBotRequestBody struct {
	Caption string `json:"caption"`
	Token   string `json:"token"`
}
