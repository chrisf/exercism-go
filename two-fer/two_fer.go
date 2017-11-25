package twofer

import "fmt"

// ShareWith returns a string "One for [name], one for me."
// If the name is a blank string, it replaces it with "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
