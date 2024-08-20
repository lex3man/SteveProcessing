package main

import (
	"core/processing/api/bot"
	"core/processing/api/front"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	frontendAPI := router.Group("/face")
	botAPI := router.Group("/api/v1")

	bot.GetRoutes(botAPI)
	front.GetRoutes(frontendAPI)

	err := router.Run("0.0.0.0:3000")
	if err != nil {
		log.Fatal("Something goes wrong...")
	}
}
