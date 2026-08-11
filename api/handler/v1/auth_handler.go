package v1

import (
	"crypto/subtle"
	"net/http"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func HandleAuthLogin(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	password, ok := requestData["password"].(string)
	configuredPassword := library.Env("ADMIN_PASSWORD")
	configuredSecret := library.Env("ADMIN_SESSION_SECRET")
	if !ok || configuredPassword == "" || configuredSecret == "" || subtle.ConstantTimeCompare([]byte(password), []byte(configuredPassword)) != 1 {
		return nil, xerror.AuthJWTInvalidTokenErr(nil)
	}

	session, err := service.GetJWTToken(configuredSecret, time.Now())
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return entity.NewObjectResponse(map[string]string{"login_token": session}), nil
}

func HandleAuthSession(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]bool{"authenticated": true}), nil
}
