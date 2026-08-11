package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	QiitaProvider = "Qiita"
	qiitaBaseURL  = "https://qiita.com/api/v2"
	qiitaPerPage  = 100
)

type QiitaItem struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	URL        string    `json:"url"`
	LikesCount int64     `json:"likes_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type QiitaClient struct {
	BaseURL    string
	User       string
	HTTPClient *http.Client
}

func NewQiitaClient(httpClient *http.Client, user string) *QiitaClient {
	if httpClient == nil {
		httpClient = defaultHTTPClient()
	}
	return &QiitaClient{BaseURL: qiitaBaseURL, User: user, HTTPClient: httpClient}
}

func (c *QiitaClient) GetUserArticle_List(ctx context.Context) ([]QiitaItem, error) {
	items := make([]QiitaItem, 0)
	for page := 1; page <= 100; page++ {
		pageItems, err := c.fetchPage(ctx, page)
		if err != nil {
			return nil, err
		}
		items = append(items, pageItems...)
		if len(pageItems) < qiitaPerPage {
			break
		}
	}
	return items, nil
}

func (c *QiitaClient) fetchPage(ctx context.Context, page int) ([]QiitaItem, error) {
	endpoint, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	endpoint.Path = strings.TrimRight(endpoint.Path, "/") + "/users/" + url.PathEscape(c.User) + "/items"
	query := endpoint.Query()
	query.Set("per_page", strconv.Itoa(qiitaPerPage))
	query.Set("page", strconv.Itoa(page))
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "buildlog-qiita-import")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Qiita API returned status %s", res.Status)
	}

	var items []QiitaItem
	if err := json.NewDecoder(res.Body).Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}
