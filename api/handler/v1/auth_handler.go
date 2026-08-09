package v1

import (
	"crypto/subtle"
	"net/http"
	"os"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func HandleAdminLogin(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	password, ok := requestData["password"].(string)
	configuredPassword := os.Getenv("ADMIN_PASSWORD")
	configuredSecret := os.Getenv("ADMIN_SESSION_SECRET")
	if !ok || configuredPassword == "" || configuredSecret == "" || subtle.ConstantTimeCompare([]byte(password), []byte(configuredPassword)) != 1 {
		return nil, xerror.AuthAdminSessionInvalid()
	}

	session, err := library.CreateAdminSession(configuredSecret, time.Now())
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return entity.NewObjectResponse(map[string]string{"session": session}), nil
}

func HandleAdminSession(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	return entity.NewObjectResponse(map[string]bool{"authenticated": true}), nil
}
