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
	c.IndentedJSON(http.StatusOK, dataBase.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range dataBase.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = uuid.New().String()

	dataBase.Albums = append(dataBase.Albums, newAlbum)
	dataBase.CollectionValue += newAlbum.Price
	c.IndentedJSON(http.StatusCreated, newAlbum)
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
}
