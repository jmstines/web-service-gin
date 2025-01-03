package dataBase

import (
	"encoding/json"
	"example/web-service-gin/models"
	"log"
	"os"
)

var Albums = make(map[string]models.Album)
var CollectionValue int64 = 0

func LoadAlbums(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var albums []models.Album
	err = json.Unmarshal(data, &albums)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for _, album := range albums {
		Albums[album.ID] = album
	}
}
