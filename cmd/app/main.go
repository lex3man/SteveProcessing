package main

import (
	"core/processing/api/bot"
	"core/processing/api/front"
	"core/processing/configs"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Error while loading env file with %s", err)
	}
}

func main() {
	router := gin.Default()

	frontendAPI := router.Group("/face")
	botAPI := router.Group("/api/v1")

	bot.GetRoutes(botAPI)
	front.GetRoutes(frontendAPI)

	config := configs.GetConfig()
	serv := fmt.Sprintf("0.0.0.0:%d", config.Server.Port)

	err := router.Run(serv)
	if err != nil {
		log.Fatal("Something goes wrong...")
	}
}
