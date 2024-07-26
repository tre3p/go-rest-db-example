package main

import (
	"example/go-rest-db/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.PostAlbum)

	router.Run("localhost:8080")
}
