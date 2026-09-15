package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

type Issue struct {
	File    string
	Message string
	Fix     string
}

func main() {
	var issues []Issue

	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			issues = append(issues, Issue{
				File:    path,
				Message: fmt.Sprintf("error reading file: %v", err),
				Fix:     "check file permissions or path",
			})
			return nil
		}

		if info.IsDir() {
			// skip vendor and .git
			if info.Name() == "vendor" || info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(path) == ".go" {
			if err := validateGoFile(path); err != nil {
				issues = append(issues, Issue{
					File:    path,
					Message: err.Error(),
					Fix:     "fix the Go syntax in this file; run `go build` or `gofmt` for hints",
				})
			}
		}

		return nil
	})

	printReport(issues)

	if len(issues) > 0 {
		os.Exit(1)
	}
}

func validateGoFile(path string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	return err
}

func printReport(issues []Issue) {
	fmt.Println("👋 Hey there! Anvil Bot here.")
	fmt.Println()

	if len(issues) == 0 {
		fmt.Println("✨ Everything looks perfect! No Go syntax issues found in this PR.")
		fmt.Println("Thanks for keeping the Anvil interpreter sharp and clean. 🚀")
		return
	}

	fmt.Printf("I checked your Go changes and found %d issue(s):\n\n", len(issues))

	for _, issue := range issues {
		fmt.Printf("📄 %s\n", issue.File)
		fmt.Printf("❌ %s\n", issue.Message)
		if issue.Fix != "" {
			fmt.Printf("💡 Suggested fix: %s\n", issue.Fix)
		}
		fmt.Println()
	}

	fmt.Println("Once you’ve fixed these, push your changes and I’ll re-check everything for you. 💻")
}
