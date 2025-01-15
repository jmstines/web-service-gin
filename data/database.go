package dataBase

import (
	"encoding/json"
	"example/web-service-gin/models"
	"log"
	"os"
)

var CollectionValue int64 = 0

func LoadAlbums(filename string) map[string]models.Album {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var albums []models.Album
	err = json.Unmarshal(data, &albums)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	albumsCollection := make(map[string]models.Album)
	for _, album := range albums {
		albumsCollection[album.ID] = album
		CollectionValue += album.Price
	}

	return albumsCollection
}
