package server

import (
	"Calendar/Controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/albums", Controller.GetAlbums)
	router.POST("/albums", Controller.PostAlbums)
	router.GET("/albums/:id", Controller.GetAlbumsByID)
	router.DELETE("/albums/:id", Controller.DeleteByID)
	router.Run("localhost:8080")
}
