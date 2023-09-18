package server

import (
	"Calendar/Controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/albums", Controller.GetAlbums)
	router.POST("/albums", Controller.PostAlbums)
	router.Run("localhost:8080")
}
