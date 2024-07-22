package controllers

import (
	"log"
	"net/http"

	utils "github.com/MobasirSarkar/youtube-url-downloader.git/utils"
)

func DownloaderHandler(w http.ResponseWriter, r *http.Request) {
	url, err := GetUrl()
	if err != nil {
		log.Fatal("error")
	}
	videoId, err := utils.VideoId(url)
	if err != nil {
		log.Fatal("error getting video")
	}

	filename, err := utils.Downloader(videoId)

	if err != nil {
		http.Error(w, "error while downloading video", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	http.ServeFile(w, r, filename)

}
