package main

import (
	"core/processing/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/v1/status", utils.GetStatus)

	err := router.Run("0.0.0.0:3000")
	if err != nil {
		log.Fatal("Something goes wrong...")
	}
}
