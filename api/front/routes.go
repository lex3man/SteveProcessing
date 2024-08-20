package front

import (
	"core/processing/configs"
	"core/processing/internal/storage"
	"core/processing/internal/utils"
	"core/processing/internal/utils/bot"
	settings "core/processing/internal/utils/front"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.RouterGroup) {
	r.GET("/status", utils.GetStatus)
	r.GET("/theme", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, settings.GetTheme())
	})
	r.GET("/settings", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, settings.GetSettings())
	})

	h := bot.Handler{DB: storage.DBInit(configs.GetDBUrl())}
	r.POST("/bots/", h.AddBot)
	r.GET("/bots", h.GetBots)
	r.GET("/bots/:id", h.GetBot)
	r.DELETE("/bots/:id", h.DeleteBot)
}
