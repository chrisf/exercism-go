package space

import (
	"strings"
)

// Planet is a string, not sure why it wasn't just used as a string in the test cases
type Planet string

const earthYearInSeconds = float64(31557600)

// Age calculates how old someone is given an age in seconds and planet
//
// - Earth: orbital period 365.25 Earth days, or 31557600 seconds
// - Mercury: orbital period 0.2408467 Earth years
// - Venus: orbital period 0.61519726 Earth years
// - Mars: orbital period 1.8808158 Earth years
// - Jupiter: orbital period 11.862615 Earth years
// - Saturn: orbital period 29.447498 Earth years
// - Uranus: orbital period 84.016846 Earth years
// - Neptune: orbital period 164.79132 Earth years
func Age(seconds float64, planet Planet) float64 {
	planet = Planet(strings.ToLower(string(planet)))
	planets := map[Planet]float64{
		"mercury": 0.2408467,
		"venus":   0.61519726,
		"earth":   1,
		"mars":    1.8808158,
		"jupiter": 11.862615,
		"saturn":  29.447498,
		"uranus":  84.016846,
		"neptune": 164.69132,
		// rip pluto
	}

	multiplier, ok := planets[planet]
	if ok {
		earthYears := seconds / earthYearInSeconds
		return earthYears / multiplier
	}

	return -1
}
