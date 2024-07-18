package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UrlInput struct {
	Url string `json:"url"`
}

var storedUrlInput UrlInput

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Homepage")
}

func UrlData(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read request body", http.StatusBadRequest)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &storedUrlInput); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(body)
}

func GetUrl() (string, error) {
	return storedUrlInput.Url, nil
}
