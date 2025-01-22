package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sahilrush/src/models"
	"gorm.io/gorm"
)

type YoutubeService struct {
	DB *gorm.DB
}

// api response
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

// FetchVideos fetches the latest Youtube videos, stores them in the database, and returns a summary response
func (services *YoutubeService) FetchVideos(query string) (string, error) {
	client := &http.Client{}
	url := "https://www.googleapis.com/youtube/v3/search"

	// Get all API keys from the environment
	apiKeys := []string{
		os.Getenv("API_KEY_1"),
		os.Getenv("API_KEY_2"),
	}

	// Ensure we have at least one API key
	if len(apiKeys) == 0 {
		return "", fmt.Errorf("no API keys found in environment variables")
	}

	// Loop through the API keys
	for i := 0; i < len(apiKeys); i++ {
		apiKey := apiKeys[i]

		// Make the request with the current API key
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %v", err)
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

		// Perform the request
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error making YouTube API request: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body: %v", err)
		}

		// Log the raw API response
		fmt.Printf("Raw YouTube API Response: %s\n", string(body))

		// Check if the error code is 403 (Quota exceeded)
		var apiResp ApiResponse
		if resp.StatusCode == 403 {
			// If quota is exceeded, print an error and continue to the next API key
			fmt.Println("Quota exceeded for API key:", apiKey)
			continue
		}

		// If the response is valid, unmarshal the JSON
		if err := json.Unmarshal(body, &apiResp); err != nil {
			return "", fmt.Errorf("error unmarshalling response: %v", err)
		}

		// Check if the response contains any items
		if len(apiResp.Items) == 0 {
			return "No results found for query: " + query, nil
		}

		// Save videos to the database
		for _, item := range apiResp.Items {
			publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if err != nil {
				return "", fmt.Errorf("error parsing published date: %v", err)
			}

			video := models.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: publishedAt,
				Thumbnails:  item.Snippet.Thumbnails.High.URL,
			}

			// Don't manually set the ID—GORM will auto-generate it
			if err := services.DB.Create(&video).Error; err != nil {
				return "", fmt.Errorf("error inserting video into the database: %v", err)
			} else {
				fmt.Println("Video inserted successfully:", video)
			}
		}

		// Return a success message with a count of videos fetched
		return fmt.Sprintf("Successfully fetched and inserted %d videos for query: %s", len(apiResp.Items), query), nil
	}

	// If all keys fail, return an error
	return "", fmt.Errorf("all API keys have exceeded their quota")
}
