package service

import (
	"bytes"
	"github.com/pari-27/GoPhotos/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestInit(t *testing.T) {
	utils.StaticRootPath = "../test_static/albums"
	os.RemoveAll(utils.StaticRootPath)
}

func TestGetAllAlbums(t *testing.T) {

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	handler := http.HandlerFunc(a.getAlbums)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateAlbums(t *testing.T) {

	rr := httptest.NewRecorder()
	jsonStr := []byte(`{"name":"album5"}`)
	req, _ := http.NewRequest("POST", "/album", bytes.NewBuffer(jsonStr))
	handler := http.HandlerFunc(a.createAlbum)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}
func TestCreateAlbumsFailure(t *testing.T) {

	rr := httptest.NewRecorder()
	jsonStr := []byte(`{"name":"album5"}`)
	req, _ := http.NewRequest("POST", "/album", bytes.NewBuffer(jsonStr))
	handler := http.HandlerFunc(a.createAlbum)
	handler.ServeHTTP(rr, req)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
func TestGetAlbumImages(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/album/album5", nil)
	handler := http.HandlerFunc(a.getAlbumImages)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAlbumImagesPathError(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/album/album6", nil)
	handler := http.HandlerFunc(a.getAlbumImages)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
