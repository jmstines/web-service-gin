package controllers

import (
	dataBase "example/web-service-gin/data"
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAlbums(c *gin.Context) {
	albums := make([]models.Album, 0, len(dataBase.Albums))
	for _, album := range dataBase.Albums {
		albums = append(albums, album)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, exists := dataBase.Albums[id]
	if exists {
		c.IndentedJSON(http.StatusOK, album)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = uuid.New().String()
	dataBase.Albums[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func PutAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	_, exists := dataBase.Albums[id]
	if exists {
		updatedAlbum.ID = id
		dataBase.Albums[id] = updatedAlbum
		c.IndentedJSON(http.StatusOK, updatedAlbum)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func GetAlbumCollectionValue(c *gin.Context) {
	if dataBase.CollectionValue == 0 && len(dataBase.Albums) > 0 {
		var prices []int64
		for _, album := range dataBase.Albums {
			prices = append(prices, album.Price)
		}
		dataBase.CollectionValue = helpers.CalculateSum(prices)
	}

	value := float64(dataBase.CollectionValue) / 100
	formattedValue := fmt.Sprintf("$ %.2f", value)

	c.IndentedJSON(http.StatusOK, gin.H{"collection_value": formattedValue})
}

func RegisterAlbumRoutes(router Router) {
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.GET("/albums/value", GetAlbumCollectionValue)
	router.POST("/albums", PostAlbums)
	router.PUT("/albums/:id", PutAlbumByID)
}
