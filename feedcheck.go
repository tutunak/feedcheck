package main

import (
	"log"
	"net/http"
	"os"
)

func lastModified(url string) (string, error) {
	// Check the Last-Modified header
	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return resp.Header.Get("Last-Modified"), nil
}
func main() {
	feed := os.Getenv("RSS_FEED")
	if feed == "" {
		log.Fatal("RSS_FEED environment variable is not set")
	}
	modificationTime, err := lastModified(feed)
	if err != nil {
		log.Fatal(err)
	}
	if modificationTime != "" {
		log.Printf("Last-Modified: %s", modificationTime)
	}
}
