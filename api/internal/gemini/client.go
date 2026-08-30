package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kikudesuyo/buildlog/api/internal/review"
)

const defaultEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent"

type Client struct {
	APIKey     string
	Model      string
	Endpoint   string
	HTTPClient *http.Client
}

type request struct {
	Contents         []content        `json:"contents"`
	GenerationConfig generationConfig `json:"generationConfig"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type generationConfig struct {
	ResponseMimeType string         `json:"responseMimeType"`
	ResponseSchema   map[string]any `json:"responseSchema"`
}

type response struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (c Client) Review(ctx context.Context, prompt string) (review.Result, error) {
	if strings.TrimSpace(c.APIKey) == "" {
		return review.Result{}, fmt.Errorf("GEMINI_API_KEY is required")
	}
	model := c.Model
	if model == "" {
		model = "gemini-2.5-flash"
	}
	endpoint := c.Endpoint
	if endpoint == "" {
		endpoint = fmt.Sprintf(defaultEndpoint, model)
	}
	body, err := json.Marshal(request{
		Contents: []content{{Parts: []part{{Text: prompt}}}},
		GenerationConfig: generationConfig{
			ResponseMimeType: "application/json",
			ResponseSchema:   schema(),
		},
	})
	if err != nil {
		return review.Result{}, fmt.Errorf("marshal Gemini request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return review.Result{}, fmt.Errorf("create Gemini request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", c.APIKey)
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return review.Result{}, fmt.Errorf("call Gemini API: %w", err)
	}
	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return review.Result{}, fmt.Errorf("read Gemini response: %w", err)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return review.Result{}, fmt.Errorf("Gemini API returned %s: %s", res.Status, strings.TrimSpace(string(responseBody)))
	}
	var decoded response
	if err := json.Unmarshal(responseBody, &decoded); err != nil {
		return review.Result{}, fmt.Errorf("decode Gemini response: %w", err)
	}
	if decoded.Error != nil {
		return review.Result{}, fmt.Errorf("Gemini API error: %s", decoded.Error.Message)
	}
	if len(decoded.Candidates) == 0 || len(decoded.Candidates[0].Content.Parts) == 0 {
		return review.Result{}, fmt.Errorf("Gemini response contains no text candidate")
	}
	return review.Parse([]byte(decoded.Candidates[0].Content.Parts[0].Text))
}

func schema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"score":       map[string]any{"type": "integer", "minimum": 0, "maximum": 100},
			"verdict":     map[string]any{"type": "string", "enum": []string{"PASS", "NEEDS_IMPROVEMENT", "FAIL"}},
			"summary":     map[string]any{"type": "string"},
			"issues":      map[string]any{"type": "array", "items": map[string]any{"type": "object", "properties": map[string]any{"severity": map[string]any{"type": "string", "enum": []string{"critical", "major", "minor"}}, "file": map[string]any{"type": "string"}, "line": map[string]any{"type": "integer", "minimum": 0}, "message": map[string]any{"type": "string"}}, "required": []string{"severity", "file", "line", "message"}}},
			"suggestions": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
		},
		"required": []string{"score", "verdict", "summary", "issues", "suggestions"},
	}
}
