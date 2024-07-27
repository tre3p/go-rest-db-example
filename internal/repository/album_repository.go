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

func (r *AlbumRepository) FindById(id int64) (models.Album, error) {
	var alb models.Album

	row := r.db.QueryRow("SELECT * FROM album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("No album found for id %d", id)
		}
		return alb, fmt.Errorf("albumsById: %d: %v", id, err)
	}

	return alb, nil
}

func (r *AlbumRepository) InsertAlbum(album *models.Album) error {
	_, err := r.db.Exec("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", album.Title, album.Artist, album.Price)
	if err != nil {
		fmt.Println("Error while adding album", err)
		return err
	}

	return nil
}
