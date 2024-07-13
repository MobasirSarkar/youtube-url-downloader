package utils

import (
	"fmt"
	"net/url"
)

func VideoId(input string) (string, error) {

	urlInput, err := url.Parse(input)

	if urlInput.Host != "www.youtube.com" {
		return "Url is not from youtube", err
	}

	if err != nil {
		return "", err
	}

	queries := urlInput.Query()
	for key, value := range queries {
		if key == "v" {
			return value[0], nil
		}
	}
	return "", fmt.Errorf("No video Id detectable")
}
