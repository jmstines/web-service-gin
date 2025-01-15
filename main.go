package main

import (
	"example/web-service-gin/controllers"
	dataBase "example/web-service-gin/data"
	"example/web-service-gin/repos"
	"example/web-service-gin/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router := controllers.NewGinRouter(engine)

	registerControllers(router)

	engine.Run("localhost:8080")
}

func registerControllers(router *controllers.GinRouter) {
	repo := repos.NewAlbumRepository()

	env := os.Getenv("ENV")
	if env == "local" {
		// Not Ideal, I couldn't find a different way to load example data for local testing.
		// .Albums should most likely be private and have a method to load the data.
		repo.Albums = dataBase.LoadAlbums("data/albums.json")
	}

	service := services.NewAlbumService(repo)
	albumController := controllers.NewAlbumController(service)
	albumController.RegisterAlbumRoutes(router)
}
