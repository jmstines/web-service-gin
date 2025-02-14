package main

import (
	"example/web-service-gin/controllers"
	dataBase "example/web-service-gin/data"
	"example/web-service-gin/repos"
	"example/web-service-gin/services"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			// Provide the dependencies for the application.
			repos.NewAlbumRepository,
			services.NewAlbumService,
			controllers.NewAlbumController,
			gin.Default,
		),
		fx.Invoke(registerRoutes),
	)

	app.Run()
}

func registerRoutes(
	albumController *controllers.AlbumControllerImp,
	engine *gin.Engine,
) {
	router := controllers.NewGinRouter(engine)

	albumController.RegisterAlbumRoutes(router)

	env := os.Getenv("ENV")
	if env == "local" {
		// Not Ideal, I couldn't find a different way to load example data for local testing.
		// .Albums should most likely be private and have a method to load the data.
		albumController.ImportDatabase(dataBase.LoadAlbums("data/albums.json"))
	}

	engine.Run("localhost:8080")
}
