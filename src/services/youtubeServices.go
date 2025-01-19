package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

// FetchVideos fetches the latest Youtube videos and stores them
func (services *YoutubeService) FetchVideos(query string) {
	client := &http.Client{}
	url := "https://www.googleapis.com/youtube/v3/search"
	apiKeyStr := os.Getenv("YOUTUBE_API_KEY")

	apiKeys := strings.Split(apiKeyStr, ",")
	if len(apiKeys) == 0 {
		log.Fatalf("No api key are found")
	}

	keyIndex := 0

	for {

		//get the current api key and iterarte thorugh it
		apiKey := apiKeys[keyIndex]
		keyIndex = (keyIndex + 1) % len(apiKeys)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("failed to connect to database %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		//adding query parameter
		q := req.URL.Query()
		q.Add("part", "snippet")
		q.Add("q", query)
		q.Add("type", "video")
		q.Add("order", "date")
		q.Add("key", apiKey)
		q.Add("maxResults", "5")
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making youtube api request:", err)
			time.Sleep(10 * time.Second)
			continue
		}
		defer resp.Body.Close()

		//parse the response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error reading response Body", err)
			time.Sleep(10 * time.Second)
			continue
		}

		var apiResp ApiResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			fmt.Println("error unmarshalling response", err)
			continue
		}

		//save videos to database
		for _, item := range apiResp.Items {
			publishedAt, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			videos := models.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: publishedAt,
				Thumbnails:  item.Snippet.Thumbnails.High.URL,
			}
			services.DB.Create(&videos)
		}
		time.Sleep(10 * time.Second)
	}
}
