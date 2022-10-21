package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pari-27/GoPhotos/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (a *App) getAlbumImages(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	albumName := vars["name"]
	var output []string
	pathToAlbum := fmt.Sprintf("%s/%s", utils.StaticRootPath, albumName)
	if _, err := os.Stat(pathToAlbum); !os.IsNotExist(err) {
		err := filepath.Walk(pathToAlbum, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
			if !info.IsDir() {
				output = append(output, path)
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	utils.SendAsJson(w, map[string]interface{}{"album": albumName, "images": output}, http.StatusOK)

}
func (a *App) addPhoto(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	albumName := r.Form.Get("albumName")
	formdata := r.MultipartForm

	//get the *fileheaders
	files := formdata.File["multiplePhotos"]

	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		dst, err := os.Create(fmt.Sprintf("%s/%s/%s", utils.StaticRootPath, albumName, files[i].Filename))

		defer dst.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(dst, file)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		log.Println(w, "Files uploaded successfully : ")
		log.Println(w, files[i].Filename+"\n")

	}
	utils.SendAsJson(w, "images added successfully", http.StatusOK)

}

func (a *App) removePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumName := vars["name"]
	picName := vars["picName"]
	err := os.Remove(fmt.Sprintf("%s/%s/%s", utils.StaticRootPath, albumName, picName))
	if err != nil {
		log.Fatal(err)
	}
	utils.SendAsJson(w, "image deleted successfully", http.StatusOK)
}
