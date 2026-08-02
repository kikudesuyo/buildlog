package v1

import (
	"encoding/json"
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetProfile はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetProfile(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	profile, err := service.GetProfile(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(profile), nil
}

// HandleUpdateProfile はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleUpdateProfile(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	var req entity.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Name == "" || req.ContactEmail == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.UpdateProfile(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}
