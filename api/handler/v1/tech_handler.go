package v1

import (
	"net/http"
	"strconv"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

func parsePaginationValue(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(value)
	if err != nil || n < 0 {
		return 0, http.ErrBodyNotAllowed
	}
	return n, nil
}

// HandleGetTech_List は外部記事の一覧を返します。
func HandleGetTech_List(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	query := r.URL.Query()
	if query.Get("all") == "true" {
		return nil, http.ErrBodyNotAllowed
	}
	offset, err := parsePaginationValue(query.Get("offset"))
	if err != nil {
		return nil, err
	}
	limit, err := parsePaginationValue(query.Get("limit"))
	if err != nil {
		return nil, err
	}
	if limit > 0 {
		limit++
	}
	order := query.Get("order")
	if order != "asc" {
		order = "desc"
	}
	techList, err := service.GetTech_List(r.Context(), false, offset, limit, order, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.New_ListResponse(techList), nil
}
