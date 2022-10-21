package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func Ok(rw http.ResponseWriter, message string) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, message)
}

func SendError(rw http.ResponseWriter, message string, code int) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	log.Printf("%v: %v %v", f.Name(), code, message)
	http.Error(rw, message, code)
}

func SendAsJson(rw http.ResponseWriter, object interface{}, code int) {
	result, err := json.Marshal(object)
	if err != nil {
		SendError(rw, "error on json marshaling: "+err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(code)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, string(result))
}

type Message struct {
	Message string
	Status  int
}
