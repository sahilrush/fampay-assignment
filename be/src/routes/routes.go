package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilrush/src/controllers"
)

// SetupRoutes sets up the routes for the API
func SetupRoutes(r *gin.Engine, vc *controllers.VideoController) {
	r.GET("/videos", vc.GetVideos)
}
