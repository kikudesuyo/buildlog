package gemini

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReviewUsesStructuredJSONOutput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-goog-api-key") != "test-key" {
			t.Errorf("missing API key header")
		}
		var requestBody map[string]any
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			t.Fatal(err)
		}
		config := requestBody["generationConfig"].(map[string]any)
		if config["responseMimeType"] != "application/json" {
			t.Errorf("unexpected response mime type: %v", config["responseMimeType"])
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"candidates":[{"content":{"parts":[{"text":"{\"score\":90,\"verdict\":\"PASS\",\"summary\":\"ok\",\"issues\":[],\"suggestions\":[]}"}]}}]}`))
	}))
	defer server.Close()

	result, err := (Client{APIKey: "test-key", Endpoint: server.URL}).Review(context.Background(), "review")
	if err != nil {
		t.Fatal(err)
	}
	if result.Score != 90 || !result.Passed() {
		t.Fatalf("unexpected result: %+v", result)
	}
}
