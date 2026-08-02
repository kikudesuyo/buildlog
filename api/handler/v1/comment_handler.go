package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

// HandleGetCommentList はHTTPリクエストを受け取り、対応する処理結果を返します。
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

// HandleDeleteComment は管理者によるコメント削除を処理します。
func HandleDeleteComment(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if os.Getenv("ADMIN_API_TOKEN") == "" || r.Header.Get("X-Admin-Token") != os.Getenv("ADMIN_API_TOKEN") {
		return nil, xerror.AuthGeneralErr(errors.New("comment management authorization failed"))
	}
	commentID, err := strconv.ParseInt(chi.URLParam(r, "commentID"), 10, 64)
	if err != nil || commentID <= 0 {
		return nil, xerror.ClientValidationErr(errors.New("invalid comment id"))
	}
	if err := service.DeleteComment(r.Context(), commentID); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}

// HandleCreateComment はHTTPリクエストを受け取り、対応する処理結果を返します。
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

// parseCommentPostID はこの処理に必要な内部処理を実行します。
func parseCommentPostID(r *http.Request) (int64, error) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || postID <= 0 {
		return 0, errors.New("invalid post id")
	}
	return postID, nil
}
