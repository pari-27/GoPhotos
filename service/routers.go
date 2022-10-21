package service

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouters(app *App) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", pong).Methods(http.MethodGet)
	r.HandleFunc("/albums", app.getAlbums).Methods(http.MethodGet)
	r.HandleFunc("/album/{name}/images", app.getAlbumImages).Methods(http.MethodGet)
	r.HandleFunc("/album", app.createAlbum).Methods(http.MethodPost)
	r.HandleFunc("/upload/image", app.addPhoto).Methods(http.MethodPost)
	r.HandleFunc("/album/{name}", app.deleteAlbum).Methods(http.MethodDelete)
	r.HandleFunc("/album/{name}/image/{picName}", app.removePhoto).Methods(http.MethodDelete)
	r.Use(RequestLoggerMiddleware(r))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	return r
}

func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				log.Printf(
					"[%s] %s %s %s",
					req.Method,
					req.Host,
					req.URL.Path,
					req.URL.RawQuery,
				)
			}()
			next.ServeHTTP(w, req)
		})
	}
}
