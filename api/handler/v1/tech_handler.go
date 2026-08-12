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

// HandleGetTechList は外部記事の一覧を返します。
func HandleGetTechList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
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
	order := query.Get("order")
	if order != "asc" {
		order = "desc"
	}
	techList, hasMore, err := service.GetTechList(r.Context(), false, offset, limit, order, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewListResponseWithHasMore(techList, hasMore), nil
}
