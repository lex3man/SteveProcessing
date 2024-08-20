package bot

import (
	"core/processing/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.RouterGroup) {
	r.GET("/status", utils.GetStatus)
}
