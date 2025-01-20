package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sahilrush/src/config"
	"github.com/sahilrush/src/controllers"
	"github.com/sahilrush/src/models"
	"github.com/sahilrush/src/routes"
	"github.com/sahilrush/src/services"
)

func main() {
	config.InitDB()

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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	videoController := &controllers.VideoController{DB: config.DB}

	routes.SetupRoutes(r, videoController)

	r.Run(":8080")
}
