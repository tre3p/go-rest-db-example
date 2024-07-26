package repository

import (
	"database/sql"
	"example/go-rest-db/internal/models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) FindAll() []models.Album {
	return make([]models.Album, 0)
}

func (r *AlbumRepository) FindById(id string) models.Album {
	return models.Album{}
}

func (r *AlbumRepository) InsertAlbum(album *models.Album) {

}
