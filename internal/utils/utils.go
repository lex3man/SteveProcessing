package utils

import (
	"core/processing/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	stat := models.Status{
		ID:    "001",
		Time:  time.Now().String(),
		Core:  "Ok",
		Bot:   "Ok",
		Admin: "Ok",
	}
	c.IndentedJSON(http.StatusOK, stat)
}
