package library

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/xconst"
)

// CtxSetJWTError はJWTの検証結果をcontextへ保存します。
// 公開APIではエラーを無視し、認証必須APIだけが利用します。
func CtxSetJWTError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, xconst.ContextKeyJWTAuth{}, err)
}

func CtxGetJWTError(ctx context.Context) (error, bool) {
	v, ok := ctx.Value(xconst.ContextKeyJWTAuth{}).(error)
	return v, ok
}

func CtxSetJWTAuthenticated(ctx context.Context, authenticated bool) context.Context {
	return context.WithValue(ctx, xconst.ContextKeyJWTAuthenticated{}, authenticated)
}

func CtxIsJWTAuthenticated(ctx context.Context) bool {
	v, _ := ctx.Value(xconst.ContextKeyJWTAuthenticated{}).(bool)
	return v
}
