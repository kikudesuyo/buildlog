package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler.ValidateRequestWithAuth(r); err != nil {
			entity.NewErrorResponse(err).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// JWTAuthToCtx はJWTを検証し、結果をcontextへ保存します。
// 公開APIを壊さないよう、このmiddleware自体は認証エラーで処理を止めません。
func JWTAuthToCtx() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if strings.HasPrefix(token, "Bearer ") {
				token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
			} else {
				token = jwtTokenValue(r)
			}

			ctx := library.CtxSetJWTAuthenticated(r.Context(), false)
			if token == "" || !service.ValidateJWTToken(token, os.Getenv("ADMIN_SESSION_SECRET")) {
				ctx = library.CtxSetJWTError(ctx, authRequiredError())
			} else {
				ctx = library.CtxSetJWTAuthenticated(ctx, true)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireAuthForAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("all") != "true" {
			next.ServeHTTP(w, r)
			return
		}
		RequireAuth(next).ServeHTTP(w, r)
	})
}

func jwtTokenValue(r *http.Request) string {
	value, err := r.Cookie(service.JWTCookie)
	if err != nil {
		return ""
	}
	return value.Value
}

func authRequiredError() error {
	return xerror.AuthJWTInvalidTokenErr(nil)
}
