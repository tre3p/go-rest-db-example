package main

import (
	"example/go-rest-db/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbum)

	router.Run("localhost:8080")
}

func readConfig() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
}

func initDbConnection() {

}
