package middleware

import (
	"net/http"
	"os"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !library.IsValidAdminSession(adminSessionValue(r), os.Getenv("ADMIN_SESSION_SECRET")) {
			entity.NewErrorResponse(authRequiredError()).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireAdminForAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("all") != "true" {
			next.ServeHTTP(w, r)
			return
		}
		RequireAdmin(next).ServeHTTP(w, r)
	})
}

func adminSessionValue(r *http.Request) string {
	value, err := r.Cookie(library.AdminSessionCookie)
	if err != nil {
		return ""
	}
	return value.Value
}

func authRequiredError() error {
	return xerror.AuthAdminSessionInvalid()
}
