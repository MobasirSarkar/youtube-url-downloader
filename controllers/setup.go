package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Print(w, "Welcome to the Home page")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func Init() {

	router := mux.NewRouter()
	setupRoutes()

	port := ":8081"
	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("Server is not running", err)
	} else {
		fmt.Printf("Server is running at port %s", port)
	}

}
