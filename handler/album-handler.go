package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerit/web-service-gin/model"
)

func GetAlbums(c *gin.Context) {
	// gin.Context carries request details, validates and serializes JSON, and more.
	// Call Context.IndentedJSON to serialize the struct into JSON
	c.IndentedJSON(http.StatusOK, model.Albums)
}
