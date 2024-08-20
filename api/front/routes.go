package front

import (
	"core/processing/internal/utils"
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
}
