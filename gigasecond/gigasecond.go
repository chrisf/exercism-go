// Package gigasecond provides a single method that adds a gigasecond to a time value
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond to a time value
func AddGigasecond(t time.Time) time.Time {
	gs := math.Pow(10, 9)
	return t.Add(time.Second * time.Duration(gs))
}
