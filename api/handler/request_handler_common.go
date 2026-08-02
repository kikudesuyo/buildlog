package handler

import "net/http"

// handleRequestBefore はこの処理に必要な内部処理を実行します。
func handleRequestBefore(r *http.Request) (*http.Request, error) {
	return r, nil
}

// handleRequestAfter はこの処理に必要な内部処理を実行します。
func handleRequestAfter(r *http.Request, responseData http.Handler) error {
	return nil
}
