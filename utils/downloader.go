package utils

import (
	"io"
	"os"
)

func Downloader(videoId string) (string, error) {

	filename := videoId + ".mp4"

	outfile, err := os.Create(filename)

	if err != nil {
		return "", err
	}

	defer outfile.Close()

	_, err = io.WriteString(outfile, "Video data"+videoId)

	if err != nil {
		return "", nil
	}

	return filename, err

}
