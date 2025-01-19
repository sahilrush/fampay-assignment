package services

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (services *YoutubeService) FetchVideos(query string) {
	client := &http.Client{}
	url := "https://www.googleapis.com/youtube/v3/search"

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("failed to connect to database %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
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
			fmt.Println("Error ")
		}
	}
}
