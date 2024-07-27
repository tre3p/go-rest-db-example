package handlers

import (
	"example/go-rest-db/internal/models"
	"example/go-rest-db/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (ah *AlbumHandler) GetAlbumById(c *gin.Context) {
	id_int, _ := strconv.Atoi(c.Param("id"))
	id := int64(id_int)

	album, err := ah.albumService.GetAlbum(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.Error{err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (ah *AlbumHandler) PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Error while binding incoming JSON to struct")
		return
	}

	err := ah.albumService.AddAlbum(&newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, models.Error{err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
