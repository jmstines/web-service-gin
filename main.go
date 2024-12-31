package main

import (
	"example/web-service-gin/controllers"
	dataBase "example/web-service-gin/data"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	env := os.Getenv("ENV")
	if env == "local" {
		dataBase.LoadAlbums("data/albums.json")
	}

	engine := gin.Default()
	router := controllers.NewGinRouter(engine)
	controllers.RegisterAlbumRoutes(router)

	engine.Run("localhost:8080")
}
