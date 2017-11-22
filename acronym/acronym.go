// Package acronym converts a phrase to an acronym
package acronym

import "strings"

// Abbreviate converts a phrase to an acronym
//
// Examples:
//
// Abbreviate("Portable Network Graphics")  => "PNG"
// Abbreviate("First In, First Out")        => "FIFO"
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	parts := strings.Split(s, " ")
	var abbrv string

	for _, part := range parts {
		abbrv += strings.ToUpper(string(part[0]))
	}

	return abbrv
}
