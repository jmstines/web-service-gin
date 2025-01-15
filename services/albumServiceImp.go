package services

import (
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"example/web-service-gin/repos"
)

type AlbumServiceImp struct {
	repo repos.AlbumRepository
}

func NewAlbumService(repo repos.AlbumRepository) *AlbumServiceImp {
	return &AlbumServiceImp{repo: repo}
}

func (service *AlbumServiceImp) GetAlbumByID(id string) (*models.Album, bool) {
	return service.repo.GetAlbumByID(id)
}

func (service *AlbumServiceImp) CreateAlbum(album *models.Album) {
	service.repo.CreateAlbum(album)

	currentValue := service.repo.GetAlbumCollectionValue()

	service.repo.UpdateAlbumCollectionValue(currentValue + album.Price)
}

func (service *AlbumServiceImp) UpdateAlbum(id string, album *models.Album) (*models.Album, bool) {
	currentValue := service.repo.GetAlbumCollectionValue()

	currentAlbum, exists := service.repo.GetAlbumByID(id)
	if exists {
		service.repo.UpdateAlbumCollectionValue(currentValue - currentAlbum.Price + album.Price)

		return service.repo.UpdateAlbum(id, album)
	}

	return nil, false
}

func (service *AlbumServiceImp) GetAllAlbums() []models.Album {
	return service.repo.GetAllAlbums()
}

func (service *AlbumServiceImp) DeleteAlbum(id string) (*models.Album, bool) {
	return service.repo.DeleteAlbum(id)
}

func (service *AlbumServiceImp) GetAlbumCollectionValue() int64 {
	albums := service.repo.GetAllAlbums()
	if service.repo.GetAlbumCollectionValue() == 0 && len(albums) > 0 {
		var prices []int64
		for _, album := range albums {
			prices = append(prices, album.Price)
		}
		value := helpers.CalculateSum(prices)

		service.repo.UpdateAlbumCollectionValue(value)
	}

	return service.repo.GetAlbumCollectionValue()
}
