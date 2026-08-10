package external

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestQiitaClientGetUserArticlesGetsEveryPage(t *testing.T) {
	client := NewQiitaClient(&http.Client{Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != "/api/v2/users/kikudesuyo/items" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("per_page") != strconv.Itoa(qiitaPerPage) {
			t.Fatalf("unexpected per_page: %s", r.URL.Query().Get("per_page"))
		}
		page, _ := url.QueryUnescape(r.URL.Query().Get("page"))
		count := 1
		if page == "1" {
			count = qiitaPerPage
		}
		items := make([]QiitaItem, count)
		for i := range items {
			items[i].ID = page + "-" + strconv.Itoa(i)
		}
		body, _ := json.Marshal(items)
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Body:       io.NopCloser(strings.NewReader(string(body))),
			Header:     make(http.Header),
			Request:    r,
		}, nil
	})}, "kikudesuyo")
	client.BaseURL = "https://qiita.example/api/v2/"
	items, err := client.GetUserArticles(context.Background())
	if err != nil {
		t.Fatalf("GetUserArticles returned error: %v", err)
	}
	if len(items) != qiitaPerPage+1 {
		t.Fatalf("expected %d items, got %d", qiitaPerPage+1, len(items))
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
