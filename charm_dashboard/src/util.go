package main

import (
	"strings"
	"unicode/utf8"
)

func truncateString(input string, maxLength int) string {
	if len(input) <= maxLength {
		return input
	}

	// Determine the length of the ellipsis (assuming it's a single character)
	ellipsisLength := utf8.RuneCountInString("…")

	// Calculate the maximum length for the truncated string
	maxTruncatedLength := maxLength - ellipsisLength

	// Truncate the input string
	truncated := input[:maxTruncatedLength]

	// Ensure the truncated string ends with a complete UTF-8 character
	for len(truncated) > 0 && !utf8.ValidString(truncated) {
		truncated = truncated[:len(truncated)-1]
	}

	// Add the ellipsis to the truncated string
	truncated += "…"

	return truncated
}

func splitLine(words []string, maxLineLength int) string {
	var result []string

	// Initialize variables to keep track of line length
	currentLine := ""
	lineLength := 0

	// Iterate over each word
	for _, word := range words {
		wordLength := len(word)

		// Check if adding the word would exceed the maximum line length
		if lineLength+wordLength+1 > maxLineLength {
			// Add the current line to the result
			result = append(result, currentLine)
			// Start a new line with the current word
			currentLine = word
			lineLength = wordLength
		} else {
			// Add the word to the current line
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
			lineLength += wordLength + 1
		}
	}

	// Add the last line to the result
	result = append(result, currentLine)

	return strings.Join(result, "\n")
}

func stringInSlice(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}
