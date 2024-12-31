package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", GetAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")
}

func TestGetAlbumByID(t *testing.T) {
	router := gin.Default()
	router.GET("/albums/:id", GetAlbumByID)

	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")

	req, _ = http.NewRequest("GET", "/albums/999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "album not found")
}

func TestPostAlbums(t *testing.T) {
	router := gin.Default()
	router.POST("/albums", PostAlbums)

	newAlbum := `{"id":"4","title":"The Modern Sound of Betty Carter","artist":"Betty Carter","price":49.99}`
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBufferString(newAlbum))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "The Modern Sound of Betty Carter")
}
