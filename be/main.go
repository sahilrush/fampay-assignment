package main

import (
	"os"
	"time"

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
	//will fetch every 10 sec
	fetchInterval := 10 * time.Second
	searchQuery := os.Getenv("SEARCH_QUERY")

	// Start periodic fetch in a goroutine
	go func() {
		for {
			_, err := youtubeService.FetchVideos(searchQuery)
			if err != nil {
				// Log error but continue
				println("Error fetching videos:", err.Error())
			}
			time.Sleep(fetchInterval)
		}
	}()

	r := gin.Default()

	//cors setup
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
