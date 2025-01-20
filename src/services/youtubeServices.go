package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sahilrush/src/models"
	"gorm.io/gorm"
)

type YoutubeService struct {
	DB *gorm.DB
}

type ApiResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  struct {
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

// FetchVideos fetches the latest Youtube videos and stores them in the database
func (services *YoutubeService) FetchVideos(query string) {
	client := &http.Client{}
	url := "https://www.googleapis.com/youtube/v3/search"
	apiKey := os.Getenv("API_KEY")

	fmt.Print(apiKey)

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("failed to create request: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Adding query parameters
		q := req.URL.Query()
		q.Add("part", "snippet")
		q.Add("q", query)
		q.Add("type", "video")
		q.Add("order", "date") // You can adjust this to "viewCount" or "relevance"
		q.Add("key", apiKey)
		q.Add("maxResults", "10") // Increased maxResults for testing
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making YouTube API request:", err)
			time.Sleep(10 * time.Second)
			continue
		}
		defer resp.Body.Close()

		// Parse the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Log the raw API response
		fmt.Printf("Raw YouTube API Response: %s\n", string(body))

		var apiResp ApiResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			fmt.Println("Error unmarshalling response:", err)
			continue
		}

		// Check if the response contains any items
		if len(apiResp.Items) == 0 {
			fmt.Println("No results found for query:", query)
		}

		// Save videos to the database
		for _, item := range apiResp.Items {
			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				fmt.Printf("Error parsing published date: %v\n", err)
				continue
			}

			video := models.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: publishedAt,
				Thumbnails:  item.Snippet.Thumbnails.High.URL,
			}

			fmt.Print("Video ", video)

			// Don't manually set the ID—GORM will auto-generate it
			if err := services.DB.Create(&video).Error; err != nil {
				fmt.Println("Error inserting video into the database:", err)
			} else {
				fmt.Println("Video inserted successfully:", video)
			}

		}

		// Wait before making the next API call
		time.Sleep(10 * time.Second)
	}
}
