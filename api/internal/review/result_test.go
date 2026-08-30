package review

import "testing"

func TestParseAndMarkdown(t *testing.T) {
	result, err := Parse([]byte(`{"score":88,"verdict":"PASS","summary":"問題ありません。","issues":[{"severity":"minor","file":"api/main.go","line":12,"message":"命名を改善できます。"}],"suggestions":["テストを追加してください。"]}`))
	if err != nil {
		t.Fatal(err)
	}
	if !result.Passed() || result.Markdown() == "" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestParseRejectsInvalidResult(t *testing.T) {
	if _, err := Parse([]byte(`{"score":101,"verdict":"PASS","summary":"x"}`)); err == nil {
		t.Fatal("expected invalid score error")
	}
	if _, err := Parse([]byte(`{"score":80,"verdict":"NEEDS_IMPROVEMENT","summary":"x","issues":null,"suggestions":[]}`)); err == nil {
		t.Fatal("expected missing issues array error")
	}
}
