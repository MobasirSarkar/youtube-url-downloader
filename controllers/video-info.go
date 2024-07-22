package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	model "github.com/MobasirSarkar/youtube-url-downloader.git/models"
	utils "github.com/MobasirSarkar/youtube-url-downloader.git/utils"

	"github.com/joho/godotenv"
)

func VideoInfo(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env file")
	}
	apiKey := os.Getenv("YOUTUBE_KEY")
	userUrl, err := GetUrl()
	if err != nil {
		log.Fatal("Error getting url")
	}
	videoId, _ := utils.VideoId(userUrl)

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s", videoId, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error while making request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}

	var videoResponse model.YouTubeResponse

	err = json.Unmarshal(body, &videoResponse)
	if err != nil {
		log.Fatalf("Error while unmarshing response : %v", err)
	}

	if len(videoResponse.Items) > 0 {
		video := videoResponse.Items[0]
		fmt.Fprint(w, video)
	} else {
		fmt.Fprint(w, "No Video Data")
	}
}
