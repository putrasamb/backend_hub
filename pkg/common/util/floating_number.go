package util

import "math"

// GetFloatingDecimal returns float64 with precision
func GetFloatingDecimal(of float64, precision float64) float64 {
	multiplier := math.Pow(10, precision)
	return math.Round(of*multiplier) / multiplier
}
