package library

import (
	"net/url"
	"strings"
)

// PublicAssetURL はGCSオブジェクトIDから公開URLを生成します。
func PublicAssetURL(objectID string) string {
	if objectID == "" || strings.HasPrefix(objectID, "/") || strings.HasPrefix(objectID, "http://") || strings.HasPrefix(objectID, "https://") {
		return objectID
	}

	baseURL := strings.TrimRight(Env("GCS_PUBLIC_BASE_URL"), "/")
	if baseURL == "" {
		baseURL = "https://storage.googleapis.com"
	}
	bucket := url.PathEscape(Env("GCS_BUCKET"))
	if bucket == "" {
		return objectID
	}

	if strings.Contains(baseURL, "/download/storage/v1/b") {
		return baseURL + "/" + bucket + "/o/" + url.PathEscape(strings.Trim(objectID, "/")) + "?alt=media"
	}

	parts := strings.Split(strings.Trim(objectID, "/"), "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}

	return baseURL + "/" + bucket + "/" + strings.Join(parts, "/")
}
