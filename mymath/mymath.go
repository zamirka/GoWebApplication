package mymath

// Sqrt is a square root function
func Sqrt(x float64) float64 {
	var z float64
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
