// Package parsinglogfiles cleans up the organization's archived log files.
package parsinglogfiles

import (
	"regexp"
)

// IsValidLine returns false if a string is not valid otherwise true.
// To be considered valid a line should begin with one of the following strings:
// [TRC], [DBG], [INF], [WRN], [ERR], [FTL]
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

// SplitLogLine takes a line and returns an array of strings each of which contains a field.
// Each field is separated by any string that has a first character of "<" and a last character of ">"
// and any combination of the following characters "~", "*", "=" and "-" in between.
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords identifies how many lines have the string "password",surrounded by quotation marks.
// "password" may be in any combination of upper/lower case and have additional content.
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[\d\D\s]*?(?i)password(?-i)[\d\D\s]*?`)
	count := 0
	for _, value := range lines {
		if re.MatchString(value) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText takes a string, removes the end-of-line text and returns a "clean" string.
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName processes log lines that contain the string "User ": prefixs the line with [USR]
// followed by the user name.
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+`)
	result := []string{}
	for _, line := range lines {
		new_line := line
		if re.MatchString(line) {
			start := re.FindAllStringIndex(line, 1)[0][1]
			end := start
			for _, char := range line[start:] {
				if char == ' ' {
					break
				}
				end++
			}
			new_line = "[USR] " + line[start:end] + " " + new_line
		}
		result = append(result, new_line)
	}
	return result
}
