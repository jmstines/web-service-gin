package services

import (
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"
	"example/web-service-gin/repos"
)

type albumServiceImpl struct {
	repo repos.AlbumRepository
}

func NewAlbumService(repo repos.AlbumRepository) AlbumService {
	return &albumServiceImpl{repo: repo}
}

func (service *albumServiceImpl) ImportDatabase(albums map[string]models.Album) {
	service.repo.ImportDatabase(albums)
}

func (service *albumServiceImpl) GetAlbumByID(id string) (*models.Album, bool) {
	return service.repo.GetAlbumByID(id)
}

func (service *albumServiceImpl) CreateAlbum(album *models.Album) {
	service.repo.CreateAlbum(album)

	currentValue := service.repo.GetAlbumCollectionValue()

	service.repo.UpdateAlbumCollectionValue(currentValue + album.Price)
}

func (service *albumServiceImpl) UpdateAlbum(id string, album *models.Album) (*models.Album, bool) {
	currentValue := service.repo.GetAlbumCollectionValue()

	currentAlbum, exists := service.repo.GetAlbumByID(id)
	if exists {
		service.repo.UpdateAlbumCollectionValue(currentValue - currentAlbum.Price + album.Price)

		return service.repo.UpdateAlbum(id, album)
	}

	return nil, false
}

func (service *albumServiceImpl) GetAllAlbums() []models.Album {
	return service.repo.GetAllAlbums()
}

func (service *albumServiceImpl) DeleteAlbum(id string) (*models.Album, bool) {
	return service.repo.DeleteAlbum(id)
}

func (service *albumServiceImpl) GetAlbumCollectionValue() int64 {
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
