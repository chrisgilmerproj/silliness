package main

import "unicode/utf8"

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
