package controllers

import (
	"example/web-service-gin/models"
	"example/web-service-gin/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumControllerImp struct {
	service services.AlbumService
}

func NewAlbumController(service services.AlbumService) *AlbumControllerImp {
	return &AlbumControllerImp{service: service}
}

func (controller *AlbumControllerImp) GetAlbums(c *gin.Context) {
	albums := controller.service.GetAllAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (controller *AlbumControllerImp) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, exists := controller.service.GetAlbumByID(id)
	if exists {
		c.IndentedJSON(http.StatusOK, album)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func (controller *AlbumControllerImp) PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	controller.service.CreateAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (controller *AlbumControllerImp) PutAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var albumToUpdate models.Album

	if err := c.BindJSON(&albumToUpdate); err != nil {
		return
	}

	updatedAlbum, exists := controller.service.UpdateAlbum(id, &albumToUpdate)
	if exists {
		c.IndentedJSON(http.StatusOK, updatedAlbum)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func (controller *AlbumControllerImp) GetAlbumCollectionValue(c *gin.Context) {

	controller.service.GetAlbumCollectionValue()

	value := float64(controller.service.GetAlbumCollectionValue()) / 100
	formattedValue := fmt.Sprintf("$ %.2f", value)

	c.IndentedJSON(http.StatusOK, gin.H{"collection_value": formattedValue})
}

func (controller *AlbumControllerImp) DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	_, exists := controller.service.DeleteAlbum(id)
	if exists {

		c.IndentedJSON(http.StatusNoContent, nil)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func (controller *AlbumControllerImp) RegisterAlbumRoutes(router Router) {
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.GET("/albums/value", controller.GetAlbumCollectionValue)
	router.POST("/albums", controller.PostAlbums)
	router.PUT("/albums/:id", controller.PutAlbumByID)
	router.DELETE("/albums/:id", controller.DeleteAlbumByID)
}
