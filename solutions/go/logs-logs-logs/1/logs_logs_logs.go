// Package logs allows users to identify which application emitted a given log,
// to fix corrupted logs, and to determine if a given log line is within a certain character limit.
package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	if strings.ContainsAny(log, "❗🔍☀") {
		for _, char := range log {
			switch char {
			case '❗':
				return "recommendation"
			case '🔍':
				return "search"
			case '☀':
				return "weather"
			}
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""
	for _, char := range log {
		if char == oldRune {
			result += string(newRune)
		} else {
			result += string(char)
		}
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) > limit {
		return false
	}
	return true
}
