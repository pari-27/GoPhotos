package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pari-27/GoPhotos/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func pong(w http.ResponseWriter, r *http.Request) {

	log.Println("came till here")
	utils.SendAsJson(w, map[string]string{"ping": "pong"}, http.StatusOK)

}

func (a *App) getAlbums(w http.ResponseWriter, r *http.Request) {

	var output []string
	err := filepath.Walk(utils.StaticRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if path != utils.StaticRootPath && info.IsDir() {
			fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
			output = append(output, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	utils.SendAsJson(w, map[string]interface{}{"albums": output}, http.StatusOK)

}

func (a *App) createAlbum(w http.ResponseWriter, r *http.Request) {

	log.Println("album creation initiated")
	var jsonInput map[string]string
	err := json.NewDecoder(r.Body).Decode(&jsonInput)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumName := fmt.Sprintf("%s/%s", utils.StaticRootPath, jsonInput["name"])
	if _, err := os.Stat(albumName); os.IsNotExist(err) {

		if err := os.MkdirAll(albumName, os.ModePerm); err != nil {
			log.Println("Failed to create albums", err)
			utils.SendError(w, "Failed to create albums", http.StatusInternalServerError)
			return
		}
		utils.SendAsJson(w, utils.Message{Message: "album created", Status: http.StatusCreated}, http.StatusCreated)

	} else {
		log.Println("album already exists")
		utils.SendError(w, "album already exists. Please provide a different name", 400)
		return

	}

}

func (a *App) deleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumName := vars["name"]
	fmt.Println("name -> ", albumName)
	err := os.RemoveAll(fmt.Sprintf("%s/%s", utils.StaticRootPath, albumName))
	if err != nil {
		fmt.Println()
		log.Fatal(err)
	}
	utils.SendAsJson(w, fmt.Sprintf("photos of album %s deleted successfully", albumName), http.StatusOK)
}
