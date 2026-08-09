package library

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const AdminSessionCookie = "buildlog_admin_session"

const adminSessionIssuer = "buildlog-api"
const adminSessionSubject = "admin"

func CreateAdminSession(secret string, now time.Time) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    adminSessionIssuer,
		Subject:   adminSessionSubject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(30 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func IsValidAdminSession(value, secret string) bool {
	if value == "" || secret == "" {
		return false
	}

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(value, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}, jwt.WithIssuer(adminSessionIssuer), jwt.WithSubject(adminSessionSubject))
	return err == nil && token.Valid
}
