package handler

import "net/http"

func handleRequestBefore(r *http.Request) (*http.Request, error) {
	return r, nil
}

func handleRequestAfter(r *http.Request, responseData http.Handler) error {
	return nil
}
