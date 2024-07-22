package routes

import (
	"fmt"
	"net/http"

	"github.com/MobasirSarkar/youtube-url-downloader.git/controllers"
	"github.com/gorilla/mux"
)

const (
	port = ":8080"
)

func setupRoutes() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", controllers.Welcome)
	myRouter.HandleFunc("/home", controllers.UrlData).Methods("POST")
	myRouter.HandleFunc("/home/videoinfo", controllers.VideoInfo).Methods("GET")
	myRouter.HandleFunc("/home/download", controllers.DownloaderHandler).Methods("GET")
	if err := http.ListenAndServe(port, myRouter); err != nil {
		fmt.Println("An Error Occured - ", err)
	}

}

func Init() {
	fmt.Println("Server Started")
	fmt.Printf("http://localhost%s", port)
	setupRoutes()
}
