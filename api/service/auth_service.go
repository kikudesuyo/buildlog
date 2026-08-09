package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWTCookie = "buildlog_jwt_token"

const jwtIssuer = "buildlog-api"
const jwtSubject = "admin"

func GetJWTToken(secret string, now time.Time) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    jwtIssuer,
		Subject:   jwtSubject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(30 * 24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWTToken(value, secret string) bool {
	if value == "" || secret == "" {
		return false
	}

	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(value, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}, jwt.WithIssuer(jwtIssuer), jwt.WithSubject(jwtSubject))
	return err == nil && token.Valid
}
