package route

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouterServesMockData(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		dataType   string
		statusCode int
	}{
		{name: "posts", path: "/api/v1/posts", dataType: "list", statusCode: http.StatusOK},
		{name: "diary", path: "/api/v1/diary", dataType: "list", statusCode: http.StatusOK},
		{name: "tech", path: "/api/v1/tech", dataType: "object", statusCode: http.StatusOK},
		{name: "profile", path: "/api/v1/profile", dataType: "object", statusCode: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			res := httptest.NewRecorder()

			NewRouter().ServeHTTP(res, req)

			if res.Code != tt.statusCode {
				t.Fatalf("expected status %d, got %d", tt.statusCode, res.Code)
			}

			var body struct {
				DataType string `json:"data_type"`
			}
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatalf("decode response: %v", err)
			}
			if body.DataType != tt.dataType {
				t.Fatalf("expected data_type %q, got %q", tt.dataType, body.DataType)
			}
		})
	}
}
