package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: glamour <filename.md>")
		os.Exit(1)
	}

	markdownFile := os.Args[1]
	markdownContent, err := os.ReadFile(markdownFile)
	if err != nil {
		fmt.Println("Error reading Markdown file:", err)
		os.Exit(1)
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		fmt.Println("Error creating Glamour renderer:", err)
		os.Exit(1)
	}

	formattedOutput, err := renderer.RenderBytes(markdownContent)
	if err != nil {
		fmt.Println("Error formatting Markdown content:", err)
		os.Exit(1)
	}

	fmt.Println(string(formattedOutput))
}
