package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/middleware"
	"github.com/kikudesuyo/buildlog/api/service"
)

func TestJWTToCtx(t *testing.T) {
	t.Setenv("ADMIN_SESSION_SECRET", "test-secret")
	validToken, err := service.GetJWTToken("test-secret", time.Now())
	if err != nil {
		t.Fatalf("GetJWTToken() error = %v", err)
	}

	testCases := []struct {
		name          string
		authorization string
		wantStatus    int
	}{
		{name: "valid bearer token", authorization: "Bearer " + validToken, wantStatus: http.StatusNoContent},
		{name: "missing token", wantStatus: http.StatusUnauthorized},
		{name: "invalid token", authorization: "Bearer invalid", wantStatus: http.StatusUnauthorized},
	}

	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler.ValidateRequestWithAuth(r); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if testCase.authorization != "" {
				req.Header.Set("Authorization", testCase.authorization)
			}
			res := httptest.NewRecorder()

			middleware.JWTToCtx()(protectedHandler).ServeHTTP(res, req)

			if res.Code != testCase.wantStatus {
				t.Errorf("status = %d, want %d", res.Code, testCase.wantStatus)
			}
		})
	}
}
