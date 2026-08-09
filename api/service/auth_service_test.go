package service

import (
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {
	now := time.Now()
	token, err := GetJWTToken("test-secret", now)
	if err != nil {
		t.Fatalf("GetJWTToken() error = %v", err)
	}
	if !ValidateJWTToken(token, "test-secret") {
		t.Fatal("ValidateJWTToken() rejected a valid token")
	}
	if ValidateJWTToken(token, "wrong-secret") {
		t.Fatal("ValidateJWTToken() accepted a token signed with another secret")
	}
}

func TestJWTTokenExpiry(t *testing.T) {
	now := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	token, err := GetJWTToken("test-secret", now.Add(-31*24*time.Hour))
	if err != nil {
		t.Fatalf("GetJWTToken() error = %v", err)
	}
	if ValidateJWTToken(token, "test-secret") {
		t.Fatal("ValidateJWTToken() accepted an expired token")
	}
}
