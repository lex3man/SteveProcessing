package bot

import (
	"core/processing/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) AddBot(c *gin.Context) {
	body := models.AddBotRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		if err := c.AbortWithError(http.StatusBadRequest, err); err != nil {
			log.Fatal(err)
		}
		return
	}

	var bot models.BotConfig

	bot.Caption = body.Caption
	bot.Token = body.Token
	bot.Active = false
	bot.Status = "test"

	if result := h.DB.Create(&bot); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Fatal(err)
		}
		return
	}

	c.JSON(http.StatusOK, &bot)
}

func (h Handler) GetBots(c *gin.Context) {
	var bots []models.BotConfig

	if result := h.DB.Find(&bots); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Fatal(err)
		}
		return
	}

	c.JSON(http.StatusOK, &bots)
}

func (h Handler) GetBot(c *gin.Context) {
	id := c.Param("id")

	var bot models.BotConfig

	if result := h.DB.First(&bot, id); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Fatal(err)
		}
		return
	}

	c.JSON(http.StatusOK, &bot)
}

func (h Handler) DeleteBot(c *gin.Context) {
	id := c.Param("id")

	var bot models.BotConfig

	if result := h.DB.First(&bot, id); result.Error != nil {
		if err := c.AbortWithError(http.StatusNotFound, result.Error); err != nil {
			log.Fatal(err)
		}
		return
	}
	h.DB.Delete(&bot)
	c.JSON(http.StatusOK, &bot)
}
