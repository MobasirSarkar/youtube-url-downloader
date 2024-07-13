package main

import (
	"fmt"

	"github.com/MobasirSarkar/youtube-url-downloader.git/controllers"
	"github.com/MobasirSarkar/youtube-url-downloader.git/utils"
)

func main() {
	url := "https://www.youtube.com/watch?v=Ihg37znaiBo"
	videoId, _ := utils.VideoId(url)
	controllers.Init()
	fmt.Println(videoId)
}
