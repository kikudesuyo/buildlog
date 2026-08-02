package v1

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func HandlePostLike(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		ipAddress := getClientIP(r)
		resp, err := service.LikePost(r.Context(), db, postID, ipAddress)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleDeleteLike(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		ipAddress := getClientIP(r)
		resp, err := service.UnlikePost(r.Context(), db, postID, ipAddress)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleGetLikeStatus(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		ipAddress := getClientIP(r)
		resp, err := service.GetLikeStatus(r.Context(), db, postID, ipAddress)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}
