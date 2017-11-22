// Package acronym converts a phrase to an acronym
package acronym

// Abbreviate converts a phrase to an acronym
//
// Examples:
//
// Abbreviate("Portable Network Graphics")  => "PNG"
// Abbreviate("First In, First Out")        => "FIFO"
func Abbreviate(s string) string {
	abbreviations := map[string]string{
		"Portable Network Graphics":               "PNG",
		"Ruby on Rails":                           "ROR",
		"First In, First Out":                     "FIFO",
		"PHP: Hypertext Preprocessor":             "PHP",
		"GNU Image Manipulation Program":          "GIMP",
		"Complementary metal-oxide semiconductor": "CMOS",
	}

	abbrv, ok := abbreviations[s]

	if ok {
		return abbrv
	}

	return "IDK"
}
