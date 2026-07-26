package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kikudesuyo/buildlog/api/handler"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/diaries", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetDiaryList(db))
		})
		r.Get("/diaries/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetDiary(db))
		})
		r.Post("/diaries", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleCreateDiary(db))
		})
		r.Put("/diaries/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleUpdateDiary(db))
		})
		r.Delete("/diaries/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleDeleteDiary(db))
		})

		r.Get("/techs", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetTechList(db))
		})
		r.Get("/techs/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetTech(db))
		})
		r.Post("/techs", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleCreateTech(db))
		})
		r.Put("/techs/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleUpdateTech(db))
		})
		r.Delete("/techs/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleDeleteTech(db))
		})

		r.Get("/trash", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetDeletedPosts(db))
		})
		r.Put("/trash/{id}/restore", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleRestorePost(db))
		})
	})

	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
