package controllers

import (
	dataBase "example/web-service-gin/data"
	"example/web-service-gin/models"
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
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func RegisterAlbumRoutes(router Router) {
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", PostAlbums)
}
