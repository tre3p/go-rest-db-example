package service

import (
	"example/go-rest-db/internal/models"
	"example/go-rest-db/internal/repository"
)

type AlbumService struct {
	albumRepo *repository.AlbumRepository
}

func NewAlbumService(albumRepo *repository.AlbumRepository) *AlbumService {
	return &AlbumService{albumRepo: albumRepo}
}

func (as *AlbumService) GetAlbums() []models.Album {
	return as.albumRepo.FindAll()
}

func (as *AlbumService) GetAlbum(id string) models.Album {
	return as.albumRepo.FindById(id)
}

func (as *AlbumService) AddAlbum(album *models.Album) {
	as.albumRepo.InsertAlbum(album)
}
