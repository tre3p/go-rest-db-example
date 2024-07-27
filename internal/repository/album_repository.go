package repository

import (
	"database/sql"
	"example/go-rest-db/internal/models"
	"fmt"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) FindAll() []models.Album {
	result := make([]models.Album, 0)

	rows, err := r.db.Query("SELECT * FROM album")
	if err != nil {
		fmt.Println("Error while querying albums", err)
		return result
	}

	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			fmt.Println("Error while scanning rows", err)
			return result
		}

		result = append(result, alb)
	}

	return result
}

func (r *AlbumRepository) FindById(id string) models.Album {
	return models.Album{}
}

func (r *AlbumRepository) InsertAlbum(album *models.Album) {

}
