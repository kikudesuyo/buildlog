package route

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	service.SetDatabase(db)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/diaries", handleFunc(v1.HandleGetDiaryList))
		r.Get("/diaries/{id}", handleFunc(v1.HandleGetDiary))
		r.Post("/diaries", handleFunc(v1.HandleCreateDiary))
		r.Put("/diaries/{id}", handleFunc(v1.HandleUpdateDiary))
		r.Delete("/diaries/{id}", handleFunc(v1.HandleDeleteDiary))

		r.Get("/techs", handleFunc(v1.HandleGetTechList))
		r.Get("/techs/{id}", handleFunc(v1.HandleGetTech))
		r.Post("/techs", handleFunc(v1.HandleCreateTech))
		r.Put("/techs/{id}", handleFunc(v1.HandleUpdateTech))
		r.Delete("/techs/{id}", handleFunc(v1.HandleDeleteTech))

<<<<<<< HEAD
		r.Post("/posts/{id}/like", handleFunc(v1.HandlePostLike))
		r.Delete("/posts/{id}/like", handleFunc(v1.HandleDeleteLike))
		r.Get("/posts/{id}/like", handleFunc(v1.HandleGetLikeStatus))
		r.Get("/trash", handleFunc(v1.HandleGetDeletedPosts))
		r.Put("/trash/{id}/restore", handleFunc(v1.HandleRestorePost))
		r.Get("/apps", handleFunc(v1.HandleGetAppList))
		r.Get("/apps/{id}", handleFunc(v1.HandleGetApp))
		r.Post("/apps", handleFunc(v1.HandleCreateApp))
		r.Put("/apps/{id}", handleFunc(v1.HandleUpdateApp))
		r.Delete("/apps/{id}", handleFunc(v1.HandleDeleteApp))
		r.Get("/profile", handleFunc(v1.HandleGetProfile))
		r.Put("/profile", handleFunc(v1.HandleUpdateProfile))
		r.Get("/admin/analytics", handleFunc(v1.HandleGetAnalytics))
=======
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
>>>>>>> dd572d718214095f25371740a9f3d14f37de939b
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
