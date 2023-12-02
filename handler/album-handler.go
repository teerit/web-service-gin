package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerit/web-service-gin/model"
)

// get Albums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	// gin.Context carries request details, validates and serializes JSON, and more.
	// Call Context.IndentedJSON to serialize the struct into JSON
	c.IndentedJSON(http.StatusOK, model.Albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	model.Albums = append(model.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range model.Albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
