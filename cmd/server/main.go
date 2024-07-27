package main

import (
	"example/go-rest-db/internal/handlers"
	"example/go-rest-db/internal/migrations"
	"example/go-rest-db/internal/repository"
	"example/go-rest-db/internal/service"
	"example/go-rest-db/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initViper()
	dbConn := database.InitDbConnection(viper.GetString("database.url"))
	migrations.RunDbMigrations(dbConn)

	albumRepo := repository.NewAlbumRepository(dbConn)
	albumService := service.NewAlbumService(albumRepo)
	albumHandler := handlers.NewAlbumHandler(albumService)

	router := gin.Default()
	router.GET("/albums", albumHandler.GetAlbums)
	router.POST("/albums", albumHandler.PostAlbum)

	router.Run("localhost:8080")
}

func initViper() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error while reading viper config", err)
		os.Exit(1)
	}
}
