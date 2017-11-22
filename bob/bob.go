// Package bob contains a function that responds to your messages
package bob

import (
	"regexp"
	"strings"
)

// Hey returns responses based on what you input
//
// If you ask a question, it will return "Sure."
// If you yell (all caps), it will return "Whoa, chill out!"
// If you send whitespace or an empty string, it will return "Fine. Be that way!"
// Else, it will return "Whatever."
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if hasUppercase(remark) && isUppercase(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	if isWhitespace(remark) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func hasUppercase(s string) bool {
	match, _ := regexp.MatchString("[A-Z]+", s)
	return match
}

func isUppercase(s string) bool {
	return strings.ToUpper(s) == s
}

func isQuestion(s string) bool {
	match, _ := regexp.MatchString("^*\\?$", s)
	return match
}

func isWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}
