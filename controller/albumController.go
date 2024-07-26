package controller

import (
	"example/go-rest-db/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums = []dto.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 177.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum dto.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("Error while binding incoming JSON to struct")
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
