package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kikudesuyo/buildlog/api/handler"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/posts", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetPostList)
		})
		r.Get("/diary", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleListDiaryEntries)
		})
		r.Get("/tech", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetTechFeed)
		})
		r.Get("/profile", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetProfileData)
		})
	})

	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
