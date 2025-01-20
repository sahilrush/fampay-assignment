package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sahilrush/src/config"
	"github.com/sahilrush/src/controllers"
	"github.com/sahilrush/src/models"
	"github.com/sahilrush/src/routes"
	"github.com/sahilrush/src/services"
)

func main() {
	config.InitDB()

	// Start background fetch service
	youtubeService := services.YoutubeService{DB: config.DB}

	config.DB.AutoMigrate(&models.Video{})

	searchQuery := os.Getenv("SEARCH_QUERY")

	go func() {
		response, err := youtubeService.FetchVideos(searchQuery)
		if err != nil {
			fmt.Println("Error fetching videos:", err)
		} else {
			fmt.Println("Response from YouTube API:", response)
		}
	}()

	r := gin.Default()

	// Create video controller
	videoController := &controllers.VideoController{DB: config.DB}

	// Set up routes
	routes.SetupRoutes(r, videoController)

	r.Run(":8080")
}
