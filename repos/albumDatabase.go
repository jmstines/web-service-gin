package repos

import "example/web-service-gin/models"

type AlbumRepository interface {
	ImportDatabase(map[string]models.Album)
	GetAlbumByID(id string) (*models.Album, bool)
	CreateAlbum(album *models.Album)
	UpdateAlbum(id string, album *models.Album) (*models.Album, bool)
	GetAllAlbums() []models.Album
	DeleteAlbum(id string) (*models.Album, bool)
	GetAlbumCollectionValue() int64
	UpdateAlbumCollectionValue(value int64)
}
