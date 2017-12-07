package hamming

import "fmt"

// Distance returns the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("The DNA strands must be the same length (got %d and %d)", len(a), len(b))
	}

	diff := 0

	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
