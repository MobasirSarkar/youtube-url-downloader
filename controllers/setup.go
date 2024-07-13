package controllers

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Print(w, "Welcome to the Home page")
}

func setupRoutes() {
	http.HandleFunc("/home", homePage)
}

func Init() {
	setupRoutes()
	fmt.Print("Server is running at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("getting error ", err)
	}
}
