package main

import (
	"github.com/pari-27/GoPhotos/service"
	"log"
	"net/http"
)

func main() {

	deps := service.Init()
	router := service.InitRouters(deps)

	err := http.ListenAndServe("localhost:9001", router)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
