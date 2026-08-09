package library

import (
	"testing"
	"time"
)

func TestAdminSession(t *testing.T) {
	now := time.Date(2026, time.August, 9, 0, 0, 0, 0, time.UTC)
	token, err := CreateAdminSession("test-secret", now)
	if err != nil {
		t.Fatalf("CreateAdminSession() error = %v", err)
	}
	if !IsValidAdminSession(token, "test-secret") {
		t.Fatal("IsValidAdminSession() rejected a valid token")
	}
	if IsValidAdminSession(token, "wrong-secret") {
		t.Fatal("IsValidAdminSession() accepted a token signed with another secret")
	}
}

func TestAdminSessionExpiry(t *testing.T) {
	now := time.Date(2026, time.August, 9, 0, 0, 0, 0, time.UTC)
	token, err := CreateAdminSession("test-secret", now.Add(-31*24*time.Hour))
	if err != nil {
		t.Fatalf("CreateAdminSession() error = %v", err)
	}
	if IsValidAdminSession(token, "test-secret") {
		t.Fatal("IsValidAdminSession() accepted an expired token")
	}
}
