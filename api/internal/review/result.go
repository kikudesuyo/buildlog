package review

import (
	"encoding/json"
	"fmt"
)

const PassScore = 85

type Result struct {
	Score       int      `json:"score"`
	Verdict     string   `json:"verdict"`
	Summary     string   `json:"summary"`
	Issues      []Issue  `json:"issues"`
	Suggestions []string `json:"suggestions"`
}

type Issue struct {
	Severity string `json:"severity"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
}

func Parse(data []byte) (Result, error) {
	var result Result
	if err := json.Unmarshal(data, &result); err != nil {
		return Result{}, fmt.Errorf("parse review JSON: %w", err)
	}
	if err := result.Validate(); err != nil {
		return Result{}, err
	}
	return result, nil
}

func (r Result) Validate() error {
	if r.Score < 0 || r.Score > 100 {
		return fmt.Errorf("review score must be between 0 and 100: %d", r.Score)
	}
	if r.Verdict != "PASS" && r.Verdict != "NEEDS_IMPROVEMENT" && r.Verdict != "FAIL" {
		return fmt.Errorf("invalid review verdict: %q", r.Verdict)
	}
	if r.Summary == "" {
		return fmt.Errorf("review summary is required")
	}
	if r.Issues == nil || r.Suggestions == nil {
		return fmt.Errorf("review issues and suggestions are required arrays")
	}
	for i, issue := range r.Issues {
		if issue.Severity != "critical" && issue.Severity != "major" && issue.Severity != "minor" {
			return fmt.Errorf("invalid issue severity at index %d: %q", i, issue.Severity)
		}
		if issue.File == "" || issue.Line < 0 || issue.Message == "" {
			return fmt.Errorf("invalid issue at index %d", i)
		}
	}
	return nil
}

func (r Result) Passed() bool { return r.Score >= PassScore }

func (r Result) Markdown() string {
	var body string
	body += "<!-- gemini-reviewer -->\n"
	body += fmt.Sprintf("## Gemini PR Review: **%d/100 — %s**\n\n", r.Score, r.Verdict)
	body += r.Summary + "\n"

	if len(r.Issues) > 0 {
		body += "\n### 指摘\n"
		for _, issue := range r.Issues {
			body += fmt.Sprintf("- **%s** `%s:%d` — %s\n", issue.Severity, issue.File, issue.Line, issue.Message)
		}
	}
	if len(r.Suggestions) > 0 {
		body += "\n### 改善提案\n"
		for _, suggestion := range r.Suggestions {
			body += "- " + suggestion + "\n"
		}
	}
	body += "\n評価基準: Correctness 30 / Maintainability 20 / Test quality 20 / Security 10 / Performance 10 / Readability 10\n"
	return body
}
