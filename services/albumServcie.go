package services

import "example/web-service-gin/models"

type AlbumService interface {
	ImportDatabase(map[string]models.Album)
	GetAlbumByID(id string) (*models.Album, bool)
	CreateAlbum(album *models.Album)
	UpdateAlbum(id string, album *models.Album) (*models.Album, bool)
	GetAllAlbums() []models.Album
	DeleteAlbum(id string) (*models.Album, bool)
	GetAlbumCollectionValue() int64
}
