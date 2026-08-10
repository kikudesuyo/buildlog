package external

import "testing"

func TestParseOGP(t *testing.T) {
	metadata := parseOGP(`<html><head>
<meta content="Qiita title" property="og:title">
<meta property="og:description" content="Qiita description">
<meta name="twitter:image" content="https://example.com/image.png">
</head></html>`)
	if metadata.Title != "Qiita title" || metadata.Description != "Qiita description" || metadata.ImageURL != "https://example.com/image.png" {
		t.Fatalf("unexpected OGP metadata: %+v", metadata)
	}
}
