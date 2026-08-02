package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func HandleGetCommentList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	postID, err := parseCommentPostID(r)
	if err != nil {
		return nil, xerror.ClientValidationErr(err)
	}

	commentList, err := service.ListCommentsByPostID(r.Context(), postID)
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(commentList), nil
}

func HandleCreateComment(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	postID, err := parseCommentPostID(r)
	if err != nil {
		return nil, xerror.ClientValidationErr(err)
	}

	var req entity.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, xerror.ClientValidationErr(err)
	}
	if strings.TrimSpace(req.Content) == "" {
		return nil, xerror.ClientValidationErr(errors.New("comment content is required"))
	}
	req.Content = strings.TrimSpace(req.Content)

	comment, err := service.CreateComment(r.Context(), postID, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(comment), nil
}

func parseCommentPostID(r *http.Request) (int64, error) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || postID <= 0 {
		return 0, errors.New("invalid post id")
	}
	return postID, nil
}
