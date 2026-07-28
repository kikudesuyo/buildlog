package route

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kikudesuyo/buildlog/api/handler"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	service.InitDB(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	handle := func(h handler.ProcessFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			handler.HandleRequestAndResponse(req, w, h)
		}
	}

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/diaries", handle(v1.HandleGetDiaryList()))
		r.Get("/diaries/{id}", handle(v1.HandleGetDiary()))
		r.Post("/diaries", handle(v1.HandleCreateDiary()))
		r.Put("/diaries/{id}", handle(v1.HandleUpdateDiary()))
		r.Delete("/diaries/{id}", handle(v1.HandleDeleteDiary()))

		r.Get("/techs", handle(v1.HandleGetTechList()))
		r.Get("/techs/{id}", handle(v1.HandleGetTech()))
		r.Post("/techs", handle(v1.HandleCreateTech()))
		r.Put("/techs/{id}", handle(v1.HandleUpdateTech()))
		r.Delete("/techs/{id}", handle(v1.HandleDeleteTech()))

		r.Post("/posts/{id}/like", handle(v1.HandlePostLike()))
		r.Delete("/posts/{id}/like", handle(v1.HandleDeleteLike()))
		r.Get("/posts/{id}/like", handle(v1.HandleGetLikeStatus()))

		r.Get("/trash", handle(v1.HandleGetDeletedPosts()))
		r.Put("/trash/{id}/restore", handle(v1.HandleRestorePost()))

		r.Get("/apps", handle(v1.HandleGetAppList()))
		r.Get("/apps/{id}", handle(v1.HandleGetApp()))
		r.Post("/apps", handle(v1.HandleCreateApp()))
		r.Put("/apps/{id}", handle(v1.HandleUpdateApp()))
		r.Delete("/apps/{id}", handle(v1.HandleDeleteApp()))

		r.Get("/profile", handle(v1.HandleGetProfile()))
		r.Put("/profile", handle(v1.HandleUpdateProfile()))
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
