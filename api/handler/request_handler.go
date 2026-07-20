package handler

import (
	"net/http"
)

type ProcessFunc func(r *http.Request, requestData map[string]interface{}) (Renderer, error)

func HandleRequestAndResponse(r *http.Request, w http.ResponseWriter, processFn ProcessFunc) {
	responseData, err := handleRequest(r, processFn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseData.Render(w)
}

func handleRequest(r *http.Request, processFn ProcessFunc) (Renderer, error) {
	request, err := handleRequestBefore(r)
	if err != nil {
		return nil, err
	}

	responseData, err := processFn(request, nil)
	if afterErr := handleRequestAfter(request, responseData); err == nil {
		err = afterErr
	}
	return responseData, err
}
