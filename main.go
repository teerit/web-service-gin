package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teerit/web-service-gin/handler"
)

func main() {
	router := gin.Default()
	// GET function to associate the GET HTTP method and /albums path with
	// a handler function
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)

	router.Run("localhost:8080")
}
