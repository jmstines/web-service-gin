package repos

import (
	"example/web-service-gin/models"

	"github.com/google/uuid"
)

type albumRepositoryImpl struct {
	Albums          map[string]models.Album
	collectionValue int64
}

func NewAlbumRepository() *albumRepositoryImpl {
	return &albumRepositoryImpl{
		Albums:          make(map[string]models.Album),
		collectionValue: 0,
	}
}

func (repo *albumRepositoryImpl) GetAlbumByID(id string) (*models.Album, bool) {
	album, exists := repo.Albums[id]
	if !exists {
		return nil, false
	}
	return &album, true
}

func (repo *albumRepositoryImpl) CreateAlbum(album *models.Album) {
	album.ID = uuid.New().String()

	repo.Albums[album.ID] = *album
}

func (repo *albumRepositoryImpl) UpdateAlbum(id string, album *models.Album) (*models.Album, bool) {
	existingAlbum, exists := repo.Albums[id]
	if !exists {
		return nil, false
	}

	repo.Albums[id] = *album
	return &existingAlbum, true
}

func (repo *albumRepositoryImpl) DeleteAlbum(id string) (*models.Album, bool) {
	album, exists := repo.Albums[id]
	if !exists {
		return nil, false
	}
	delete(repo.Albums, id)
	return &album, true
}

func (repo *albumRepositoryImpl) GetAllAlbums() []models.Album {
	albums := make([]models.Album, 0, len(repo.Albums))
	for _, album := range repo.Albums {
		albums = append(albums, album)
	}
	return albums
}

func (repo *albumRepositoryImpl) UpdateAlbumCollectionValue(value int64) {
	repo.collectionValue = value
}

func (repo *albumRepositoryImpl) GetAlbumCollectionValue() int64 {
	return repo.collectionValue
}
