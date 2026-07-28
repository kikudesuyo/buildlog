package v1

import (
	"encoding/json"
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetProfile() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		profile, err := service.GetProfile(r.Context())
		if err != nil {
			return nil, err
		}
		return entity.NewObjectResponse(profile), nil
	}
}

func HandleUpdateProfile() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var req entity.UpdateProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}

		// 最低限のバリデーション
		if req.Name == "" || req.ContactEmail == "" {
			return nil, http.ErrBodyNotAllowed
		}

		resp, err := service.UpdateProfile(r.Context(), req)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}
