package util

// GetAmountFromPercentage calculate percentage amount. eg., 50% of 1000 returns 500
func GetAmountFromPercentage(percentage float64, from float64) float64 {
	return float64(percentage) / 100 * from
}
