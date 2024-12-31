package dataBase

import (
	"example/web-service-gin/models"
	"os"

	"encoding/json"
	"log"
)

var Albums []models.Album

func LoadAlbums(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &Albums)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
	}
}
