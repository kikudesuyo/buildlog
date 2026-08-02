package route

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdminAuthMiddlewareReturnsCommonUnauthorizedResponse(t *testing.T) {
	t.Setenv("ADMIN_TOKEN", "test-admin-token")

	handler := adminAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	req := httptest.NewRequest(http.MethodPost, "/api/v1/diaries", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusUnauthorized)
	}
	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", contentType)
	}
	if !strings.Contains(recorder.Body.String(), `"data_type":"error"`) {
		t.Fatalf("response body does not use the common error response: %s", recorder.Body.String())
	}
}

func TestAdminAuthMiddlewareAllowsPreflightWithoutToken(t *testing.T) {
	handler := adminAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	req := httptest.NewRequest(http.MethodOptions, "/api/v1/diaries", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusNoContent)
	}
}
