package route

import (
	"net/http"
	"os"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func adminAuthMiddleware(next http.Handler) http.Handler {
	adminToken := os.Getenv("ADMIN_TOKEN")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if adminToken == "" || authHeader == "" {
			entity.NewErrorResponse(xerror.AuthMissingHeader()).ServeHTTP(w, r)
			return
		}
		if authHeader != "Bearer "+adminToken {
			entity.NewErrorResponse(xerror.AuthInvalidHeader()).ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
