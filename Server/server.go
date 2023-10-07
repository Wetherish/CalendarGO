package server

import (
	"Calendar/Controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/albums", Controller.GetAlbums)
	router.POST("/albums", Controller.PostAlbums)
	router.GET("/albums/:id", Controller.GetAlbums)
	router.DELETE("/albums/:id", Controller.GetAlbumsByID)
	router.Run("localhost:8080")
}
