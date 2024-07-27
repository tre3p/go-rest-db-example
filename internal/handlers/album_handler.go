package handlers

import (
	"example/go-rest-db/internal/models"
	"example/go-rest-db/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlbumHandler struct {
	albumService *service.AlbumService
}

var albums = []models.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 177.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func NewAlbumHandler(albumService *service.AlbumService) *AlbumHandler {
	return &AlbumHandler{albumService: albumService}
}

func (ah *AlbumHandler) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ah.albumService.GetAlbums())
}

func (ah *AlbumHandler) GetAlbumById(id string) *models.Album {
	//album := ah.albumService.GetAlbum(id)

	return &models.Album{}
}

func (ah *AlbumHandler) PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Error while binding incoming JSON to struct")
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
