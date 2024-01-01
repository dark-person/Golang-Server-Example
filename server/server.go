package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("web/build"))
	r.Handle("/static/*", fileServer)
	r.Handle("/", fileServer)

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

	// r.Route("/gallery", func(r chi.Router) {
	// 	r.With(paginate).Get("/", ListArticles)
	// 	r.Post("/", CreateArticle)       // POST /articles
	// 	r.Get("/search", SearchArticles) // GET /articles/search

	// 	r.Route("/{articleID}", func(r chi.Router) {
	// 		r.Use(ArticleCtx)            // Load the *Article on the request context
	// 		r.Get("/", GetArticle)       // GET /articles/123
	// 		r.Put("/", UpdateArticle)    // PUT /articles/123
	// 		r.Delete("/", DeleteArticle) // DELETE /articles/123
	// 	})

	// 	// GET /articles/whats-up
	// 	r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	// })

	return r
}
