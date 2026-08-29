package library

import "testing"

func TestPublicAssetURL(t *testing.T) {
	t.Setenv("IS_PRODUCTION", "")
	t.Setenv("GCS_BUCKET", "buildlog-local")
	t.Setenv("GCS_PUBLIC_BASE_URL", "http://localhost:14443/download/storage/v1/b")

	got := PublicAssetURL("apps/whichway-icon.svg")
	want := "http://localhost:14443/download/storage/v1/b/buildlog-local/o/apps%2Fwhichway-icon.svg?alt=media"
	if got != want {
		t.Fatalf("PublicAssetURL() = %q, want %q", got, want)
	}
}

func TestPublicAssetURLKeepsAbsoluteURL(t *testing.T) {
	absoluteURL := "https://example.com/image.jpg"
	if got := PublicAssetURL(absoluteURL); got != absoluteURL {
		t.Fatalf("PublicAssetURL() = %q, want %q", got, absoluteURL)
	}
}
