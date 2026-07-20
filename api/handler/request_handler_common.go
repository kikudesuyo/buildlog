package handler

import "net/http"

func handleRequestBefore(r *http.Request) (*http.Request, error) {
	return r, nil
}

func handleRequestAfter(r *http.Request, responseData Renderer) error {
	return nil
}
