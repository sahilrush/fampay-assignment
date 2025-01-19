package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilrush/src/config"
	"github.com/sahilrush/src/controllers"
	"github.com/sahilrush/src/routes"
	"github.com/sahilrush/src/services"
)

func main() {
	config.InitDB()

	// Start background fetch service
	youtubeService := services.YoutubeService{DB: config.DB}
	go youtubeService.FetchVideos("cricket") // Example query

	r := gin.Default()

	// Create video controller
	videoController := &controllers.VideoController{DB: config.DB}

	// Set up routes
	routes.SetupRoutes(r, videoController)

	r.Run(":8080")
}
