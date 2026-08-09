package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

// JWTToCtx はJWTを検証し、結果をcontextへ保存します。
// 認証が必要かどうかの判定は各handlerが行います。
func JWTToCtx() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if strings.HasPrefix(token, "Bearer ") {
				token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
			} else {
				token = jwtTokenValue(r)
			}

			ctx := r.Context()
			if token == "" {
				ctx = library.CtxSetJWTError(ctx, xerror.AuthJWTEmptyToken())
			} else if !service.ValidateJWTToken(token, os.Getenv("ADMIN_SESSION_SECRET")) {
				ctx = library.CtxSetJWTError(ctx, xerror.AuthJWTInvalidTokenErr(nil))
			} else {
				ctx = library.CtxSetJWTAuthenticated(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func jwtTokenValue(r *http.Request) string {
	value, err := r.Cookie(service.JWTCookie)
	if err != nil {
		return ""
	}
	return value.Value
}
