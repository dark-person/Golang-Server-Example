package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static file of react
	fileServer := http.FileServer(http.Dir("web/build"))
	r.Handle("/static/*", fileServer)
	r.Handle("/", fileServer)

	// API endpoint
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("api entry: please refer to document"))
		})

		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("api get all post.."))
		})

		r.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(chi.URLParam(r, "id"))
			fmt.Fprintf(w, "api get post id: %d", id)
		})

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ping.."))
		})
	})

	return r
}
