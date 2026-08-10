package external

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type OGPMetadata struct {
	Title       string
	Description string
	ImageURL    string
}

var metaTagPattern = regexp.MustCompile(`(?is)<meta\s+[^>]*>`)
var attributePattern = regexp.MustCompile(`(?i)(property|name|content)=["']([^"']*)["']`)

func (c *QiitaClient) GetOGP(ctx context.Context, rawURL string) (OGPMetadata, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Scheme != "https" || parsed.Host != "qiita.com" {
		return OGPMetadata{}, fmt.Errorf("unsupported OGP URL")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsed.String(), nil)
	if err != nil {
		return OGPMetadata{}, err
	}
	req.Header.Set("User-Agent", "buildlog-qiita-import")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return OGPMetadata{}, err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return OGPMetadata{}, fmt.Errorf("OGP page returned status %s", res.Status)
	}

	body, err := io.ReadAll(io.LimitReader(res.Body, 2<<20))
	if err != nil {
		return OGPMetadata{}, err
	}
	return parseOGP(string(body)), nil
}

func parseOGP(body string) OGPMetadata {
	var metadata OGPMetadata
	for _, tag := range metaTagPattern.FindAllString(body, -1) {
		attrs := map[string]string{}
		for _, match := range attributePattern.FindAllStringSubmatch(tag, -1) {
			attrs[strings.ToLower(match[1])] = html.UnescapeString(strings.TrimSpace(match[2]))
		}
		switch strings.ToLower(attrs["property"]) {
		case "og:title":
			metadata.Title = attrs["content"]
		case "og:description":
			metadata.Description = attrs["content"]
		case "og:image":
			metadata.ImageURL = httpsImageURL(attrs["content"])
		}
		switch strings.ToLower(attrs["name"]) {
		case "twitter:title":
			if metadata.Title == "" {
				metadata.Title = attrs["content"]
			}
		case "twitter:description":
			if metadata.Description == "" {
				metadata.Description = attrs["content"]
			}
		case "twitter:image":
			if metadata.ImageURL == "" {
				metadata.ImageURL = httpsImageURL(attrs["content"])
			}
		}
	}
	return metadata
}

func httpsImageURL(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
		return ""
	}
	return parsed.String()
}

func defaultHTTPClient() *http.Client { return &http.Client{Timeout: 15 * time.Second} }
