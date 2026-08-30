package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kikudesuyo/buildlog/api/internal/gemini"
	"github.com/kikudesuyo/buildlog/api/internal/review"
)

func main() {
	diffPath := flag.String("diff", "diff.txt", "path to the pull request diff")
	testPath := flag.String("test-result", "test-result.txt", "path to test output")
	lintPath := flag.String("lint-result", "lint-result.txt", "path to lint output")
	flag.Parse()

	diff := readFile(*diffPath)
	testResult := readFile(*testPath)
	lintResult := readFile(*lintPath)
	result, err := (gemini.Client{APIKey: os.Getenv("GEMINI_API_KEY"), Model: os.Getenv("GEMINI_MODEL")}).Review(context.Background(), review.Prompt(review.Input{
		Diff: diff, TestResult: testResult, LintResult: lintResult,
	}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result.Markdown())
	if !result.Passed() {
		os.Exit(2)
	}
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}
