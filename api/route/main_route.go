package route

import (
	"net/http"
	"os"

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

		r.Post("/posts/{id}/like", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandlePostLike(db))
		})
		r.Delete("/posts/{id}/like", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleDeleteLike(db))
		})
		r.Get("/posts/{id}/like", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetLikeStatus(db))
		})
		r.Get("/trash", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetDeletedPosts(db))
		})
		r.Put("/trash/{id}/restore", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleRestorePost(db))
		})
		r.Get("/apps", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetAppList(db))
		})
		r.Get("/apps/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetApp(db))
		})
		r.Post("/apps", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleCreateApp(db))
		})
		r.Put("/apps/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleUpdateApp(db))
		})
		r.Delete("/apps/{id}", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleDeleteApp(db))
		})
		r.Get("/profile", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetProfile(db))
		})
		r.Put("/profile", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleUpdateProfile(db))
		})
		r.Get("/posts/history", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetPostHistory(db))
		})
		r.Get("/admin/analytics", func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, v1.HandleGetAnalytics(db))
		})
	})

	return r
}
func corsMiddleware(next http.Handler) http.Handler {
	allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigin == "" {
		allowedOrigin = "*"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
