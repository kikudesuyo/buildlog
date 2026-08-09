package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler.ValidateRequestWithAuth(r); err != nil {
			entity.NewErrorResponse(err).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// JWTAdminToCtx はJWTを検証し、結果をcontextへ保存します。
// 公開APIを壊さないよう、このmiddleware自体は認証エラーで処理を止めません。
func JWTAdminToCtx() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if strings.HasPrefix(token, "Bearer ") {
				token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
			} else {
				token = adminSessionValue(r)
			}

			ctx := library.CtxSetAdminAuthenticated(r.Context(), false)
			if token == "" {
				ctx = library.CtxSetJWTError(ctx, authRequiredError())
			} else if !library.IsValidAdminSession(token, os.Getenv("ADMIN_SESSION_SECRET")) {
				ctx = library.CtxSetJWTError(ctx, authRequiredError())
			} else {
				ctx = library.CtxSetAdminAuthenticated(ctx, true)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
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
