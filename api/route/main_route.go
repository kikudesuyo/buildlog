package route

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
	authmiddleware "github.com/kikudesuyo/buildlog/api/middleware"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

// NewRouter は値を生成します。
func NewRouter(db *gorm.DB) http.Handler {
	service.SetDatabase(db)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/login", handleFunc(v1.HandleAdminLogin))
		r.With(authmiddleware.RequireAdmin).Get("/auth/session", handleFunc(v1.HandleAdminSession))
		r.With(authmiddleware.RequireAdminForAll).Get("/diaries", handleFunc(v1.HandleGetDiaryList))
		r.Get("/diaries/{id}", handleFunc(v1.HandleGetDiary))
		r.With(authmiddleware.RequireAdmin).Post("/diaries", handleFunc(v1.HandleCreateDiary))
		r.With(authmiddleware.RequireAdmin).Put("/diaries/{id}", handleFunc(v1.HandleUpdateDiary))
		r.With(authmiddleware.RequireAdmin).Delete("/diaries/{id}", handleFunc(v1.HandleDeleteDiary))

		r.With(authmiddleware.RequireAdminForAll).Get("/techs", handleFunc(v1.HandleGetTechList))
		r.Get("/techs/{id}", handleFunc(v1.HandleGetTech))
		r.With(authmiddleware.RequireAdmin).Post("/techs", handleFunc(v1.HandleCreateTech))
		r.With(authmiddleware.RequireAdmin).Put("/techs/{id}", handleFunc(v1.HandleUpdateTech))
		r.With(authmiddleware.RequireAdmin).Delete("/techs/{id}", handleFunc(v1.HandleDeleteTech))

		r.Post("/posts/{id}/like", handleFunc(v1.HandlePostLike))
		r.Delete("/posts/{id}/like", handleFunc(v1.HandleDeleteLike))
		r.Get("/posts/{id}/like", handleFunc(v1.HandleGetLikeStatus))
		r.With(authmiddleware.RequireAdmin).Get("/trash", handleFunc(v1.HandleGetDeletedPosts))
		r.With(authmiddleware.RequireAdmin).Put("/trash/{id}/restore", handleFunc(v1.HandleRestorePost))
		r.Get("/apps", handleFunc(v1.HandleGetAppList))
		r.With(authmiddleware.RequireAdmin).Get("/apps/{id}", handleFunc(v1.HandleGetApp))
		r.With(authmiddleware.RequireAdmin).Post("/apps", handleFunc(v1.HandleCreateApp))
		r.With(authmiddleware.RequireAdmin).Put("/apps/{id}", handleFunc(v1.HandleUpdateApp))
		r.With(authmiddleware.RequireAdmin).Delete("/apps/{id}", handleFunc(v1.HandleDeleteApp))
		r.Get("/profile", handleFunc(v1.HandleGetProfile))
		r.With(authmiddleware.RequireAdmin).Put("/profile", handleFunc(v1.HandleUpdateProfile))
		r.Get("/posts/{id}/comments", handleFunc(v1.HandleGetCommentList))
		r.Post("/posts/{id}/comments", handleFunc(v1.HandleCreateComment))
		r.With(authmiddleware.RequireAdmin).Get("/posts/history", handleFunc(v1.HandleGetPostHistory(db)))
		r.With(authmiddleware.RequireAdmin).Get("/admin/analytics", handleFunc(v1.HandleGetAnalytics))
		r.With(authmiddleware.RequireAdmin).Get("/goals/current", handleFunc(v1.HandleGetCurrentGoals))
		r.With(authmiddleware.RequireAdmin).Put("/goals/current", handleFunc(v1.HandleSaveCurrentGoals))
	})

	return r
}

// corsMiddleware はこの処理に必要な内部処理を実行します。
func corsMiddleware(next http.Handler) http.Handler {
	allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigin == "" {
		allowedOrigin = "*"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if allowedOrigin != "*" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
